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
	sysRoleFieldNames          = builder.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysRoleIdPrefix = "cache:sysRole:id:"
)

type (
	SysRoleModel interface {
		Insert(data *SysRole) (sql.Result, error)
		FindOne(id int64) (*SysRole, error)
		Update(data *SysRole) error
		Delete(id int64) error
	}

	defaultSysRoleModel struct {
		sqlc.CachedConn
		table string
	}

	SysRole struct {
		Id        int64        `db:"id"`         // 编号
		Name      string       `db:"name"`       // 角色名称
		Remark    string       `db:"remark"`     // 备注
		Status    int64        `db:"status"`     // 状态  1:启用,0:禁用
		CreatedBy string       `db:"created_by"` // 创建人
		UpdatedBy string       `db:"updated_by"` // 更新人
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func NewSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleModel {
	return &defaultSysRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Insert(data *SysRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.Remark, data.Status, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultSysRoleModel) FindOne(id int64) (*SysRole, error) {
	sysRoleIdKey := fmt.Sprintf("%s%v", cacheSysRoleIdPrefix, id)
	var resp SysRole
	err := m.QueryRow(&resp, sysRoleIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
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

func (m *defaultSysRoleModel) Update(data *SysRole) error {
	sysRoleIdKey := fmt.Sprintf("%s%v", cacheSysRoleIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Remark, data.Status, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, sysRoleIdKey)
	return err
}

func (m *defaultSysRoleModel) Delete(id int64) error {

	sysRoleIdKey := fmt.Sprintf("%s%v", cacheSysRoleIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysRoleIdKey)
	return err
}

func (m *defaultSysRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysRoleIdPrefix, primary)
}

func (m *defaultSysRoleModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
	return conn.QueryRow(v, query, primary)
}
