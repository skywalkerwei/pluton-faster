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
	sysConfigFieldNames          = builder.RawFieldNames(&SysConfig{})
	sysConfigRows                = strings.Join(sysConfigFieldNames, ",")
	sysConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysConfigIdPrefix = "cache:sysConfig:id:"
)

type (
	SysConfigModel interface {
		Insert(data *SysConfig) (sql.Result, error)
		FindOne(id int64) (*SysConfig, error)
		Update(data *SysConfig) error
		Delete(id int64) error
	}

	defaultSysConfigModel struct {
		sqlc.CachedConn
		table string
	}

	SysConfig struct {
		Id          int64        `db:"id"`          // 编号
		Value       string       `db:"value"`       // 数据值
		Label       string       `db:"label"`       // 标签名
		Tp          string       `db:"tp"`          // 类型
		Description string       `db:"description"` // 描述
		Sort        int64        `db:"sort"`        // 排序（升序）
		Remarks     string       `db:"remarks"`     // 备注信息
		CreatedBy   string       `db:"created_by"`  // 创建人
		UpdatedBy   string       `db:"updated_by"`  // 更新人
		CreatedAt   time.Time    `db:"created_at"`
		UpdatedAt   time.Time    `db:"updated_at"`
		DeletedAt   sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func NewSysConfigModel(conn sqlx.SqlConn, c cache.CacheConf) SysConfigModel {
	return &defaultSysConfigModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_config`",
	}
}

func (m *defaultSysConfigModel) Insert(data *SysConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysConfigRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Value, data.Label, data.Tp, data.Description, data.Sort, data.Remarks, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultSysConfigModel) FindOne(id int64) (*SysConfig, error) {
	sysConfigIdKey := fmt.Sprintf("%s%v", cacheSysConfigIdPrefix, id)
	var resp SysConfig
	err := m.QueryRow(&resp, sysConfigIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysConfigRows, m.table)
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

func (m *defaultSysConfigModel) Update(data *SysConfig) error {
	sysConfigIdKey := fmt.Sprintf("%s%v", cacheSysConfigIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysConfigRowsWithPlaceHolder)
		return conn.Exec(query, data.Value, data.Label, data.Tp, data.Description, data.Sort, data.Remarks, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, sysConfigIdKey)
	return err
}

func (m *defaultSysConfigModel) Delete(id int64) error {

	sysConfigIdKey := fmt.Sprintf("%s%v", cacheSysConfigIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysConfigIdKey)
	return err
}

func (m *defaultSysConfigModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysConfigIdPrefix, primary)
}

func (m *defaultSysConfigModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysConfigRows, m.table)
	return conn.QueryRow(v, query, primary)
}
