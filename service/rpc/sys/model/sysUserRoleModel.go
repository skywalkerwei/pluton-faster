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
	sysUserRoleFieldNames          = builder.RawFieldNames(&SysUserRole{})
	sysUserRoleRows                = strings.Join(sysUserRoleFieldNames, ",")
	sysUserRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserRoleFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysUserRoleIdPrefix = "cache:sysUserRole:id:"
)

type (
	SysUserRoleModel interface {
		Insert(data *SysUserRole) (sql.Result, error)
		FindOne(id int64) (*SysUserRole, error)
		Update(data *SysUserRole) error
		Delete(id int64) error
	}

	defaultSysUserRoleModel struct {
		sqlc.CachedConn
		table string
	}

	SysUserRole struct {
		Id        int64     `db:"id"`         // 编号
		UserId    int64     `db:"user_id"`    // 用户ID
		RoleId    int64     `db:"role_id"`    // 角色ID
		CreatedBy string    `db:"created_by"` // 创建人
		UpdatedBy string    `db:"updated_by"` // 更新人
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func NewSysUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserRoleModel {
	return &defaultSysUserRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_user_role`",
	}
}

func (m *defaultSysUserRoleModel) Insert(data *SysUserRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysUserRoleRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserId, data.RoleId, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt)

	return ret, err
}

func (m *defaultSysUserRoleModel) FindOne(id int64) (*SysUserRole, error) {
	sysUserRoleIdKey := fmt.Sprintf("%s%v", cacheSysUserRoleIdPrefix, id)
	var resp SysUserRole
	err := m.QueryRow(&resp, sysUserRoleIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRoleRows, m.table)
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

func (m *defaultSysUserRoleModel) Update(data *SysUserRole) error {
	sysUserRoleIdKey := fmt.Sprintf("%s%v", cacheSysUserRoleIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRoleRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.RoleId, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.Id)
	}, sysUserRoleIdKey)
	return err
}

func (m *defaultSysUserRoleModel) Delete(id int64) error {

	sysUserRoleIdKey := fmt.Sprintf("%s%v", cacheSysUserRoleIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysUserRoleIdKey)
	return err
}

func (m *defaultSysUserRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysUserRoleIdPrefix, primary)
}

func (m *defaultSysUserRoleModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRoleRows, m.table)
	return conn.QueryRow(v, query, primary)
}
