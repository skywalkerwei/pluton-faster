package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysLoginLogFieldNames          = builder.RawFieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysLoginLogIdPrefix = "cache:sysLoginLog:id:"
)

type (
	SysLoginLogModel interface {
		Insert(data *SysLoginLog) (sql.Result, error)
		FindOne(id int64) (*SysLoginLog, error)
		Update(data *SysLoginLog) error
		Delete(id int64) error
	}

	defaultSysLoginLogModel struct {
		sqlc.CachedConn
		table string
	}

	SysLoginLog struct {
		Id        int64     `db:"id"`         // 编号
		UserName  string    `db:"user_name"`  // 用户名
		Status    string    `db:"status"`     // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
		Ip        string    `db:"ip"`         // IP地址
		CreatedBy string    `db:"created_by"` // 创建人
		UpdatedBy string    `db:"updated_by"` // 更新人
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func NewSysLoginLogModel(conn sqlx.SqlConn, c cache.CacheConf) SysLoginLogModel {
	return &defaultSysLoginLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_login_log`",
	}
}

func (m *defaultSysLoginLogModel) Insert(data *SysLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserName, data.Status, data.Ip, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt)

	return ret, err
}

func (m *defaultSysLoginLogModel) FindOne(id int64) (*SysLoginLog, error) {
	sysLoginLogIdKey := fmt.Sprintf("%s%v", cacheSysLoginLogIdPrefix, id)
	var resp SysLoginLog
	err := m.QueryRow(&resp, sysLoginLogIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysLoginLogModel) Update(data *SysLoginLog) error {
	sysLoginLogIdKey := fmt.Sprintf("%s%v", cacheSysLoginLogIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLoginLogRowsWithPlaceHolder)
		return conn.Exec(query, data.UserName, data.Status, data.Ip, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.Id)
	}, sysLoginLogIdKey)
	return err
}

func (m *defaultSysLoginLogModel) Delete(id int64) error {

	sysLoginLogIdKey := fmt.Sprintf("%s%v", cacheSysLoginLogIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysLoginLogIdKey)
	return err
}

func (m *defaultSysLoginLogModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysLoginLogIdPrefix, primary)
}

func (m *defaultSysLoginLogModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
	return conn.QueryRow(v, query, primary)
}
