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
	sysLogFieldNames          = builder.RawFieldNames(&SysLog{})
	sysLogRows                = strings.Join(sysLogFieldNames, ",")
	sysLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysLogIdPrefix = "cache:sysLog:id:"
)

type (
	SysLogModel interface {
		Insert(data *SysLog) (sql.Result, error)
		FindOne(id int64) (*SysLog, error)
		Update(data *SysLog) error
		Delete(id int64) error
		FindAll(int64, int64) ([]*SysLog, error)
		Count() (int64, error)
	}

	defaultSysLogModel struct {
		sqlc.CachedConn
		table string
	}

	SysLog struct {
		Id        int64     `db:"id"`         // 编号
		UserName  string    `db:"user_name"`  // 用户名
		Operation string    `db:"operation"`  // 用户操作
		Method    string    `db:"method"`     // 请求方法
		Params    string    `db:"params"`     // 请求参数
		Time      int64     `db:"time"`       // 执行时长(毫秒)
		Ip        string    `db:"ip"`         // IP地址
		CreatedBy string    `db:"created_by"` // 创建人
		UpdatedBy string    `db:"updated_by"` // 更新人
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func NewSysLogModel(conn sqlx.SqlConn, c cache.CacheConf) SysLogModel {
	return &defaultSysLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_log`",
	}
}

func (m *defaultSysLogModel) Insert(data *SysLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysLogRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt)

	return ret, err
}

func (m *defaultSysLogModel) FindOne(id int64) (*SysLog, error) {
	sysLogIdKey := fmt.Sprintf("%s%v", cacheSysLogIdPrefix, id)
	var resp SysLog
	err := m.QueryRow(&resp, sysLogIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
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

func (m *defaultSysLogModel) Update(data *SysLog) error {
	sysLogIdKey := fmt.Sprintf("%s%v", cacheSysLogIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLogRowsWithPlaceHolder)
		return conn.Exec(query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.Id)
	}, sysLogIdKey)
	return err
}

func (m *defaultSysLogModel) Delete(id int64) error {

	sysLogIdKey := fmt.Sprintf("%s%v", cacheSysLogIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysLogIdKey)
	return err
}

func (m *defaultSysLogModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysLogIdPrefix, primary)
}

func (m *defaultSysLogModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultSysLogModel) FindAll(Current int64, PageSize int64) ([]*SysLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysLogRows, m.table)
	var resp []*SysLog
	//err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	err := m.QueryRowsNoCache(&resp, query, (Current-1)*PageSize, PageSize)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysLogModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.QueryRowNoCache(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
