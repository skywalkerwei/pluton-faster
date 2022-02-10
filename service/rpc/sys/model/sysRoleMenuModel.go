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
	sysRoleMenuFieldNames          = builder.RawFieldNames(&SysRoleMenu{})
	sysRoleMenuRows                = strings.Join(sysRoleMenuFieldNames, ",")
	sysRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysRoleMenuIdPrefix = "cache:sysRoleMenu:id:"
)

type (
	SysRoleMenuModel interface {
		Insert(data *SysRoleMenu) (sql.Result, error)
		FindOne(id int64) (*SysRoleMenu, error)
		Update(data *SysRoleMenu) error
		Delete(id int64) error
	}

	defaultSysRoleMenuModel struct {
		sqlc.CachedConn
		table string
	}

	SysRoleMenu struct {
		Id        int64     `db:"id"`         // 编号
		RoleId    int64     `db:"role_id"`    // 角色ID
		MenuId    int64     `db:"menu_id"`    // 菜单ID
		CreatedBy string    `db:"created_by"` // 创建人
		UpdatedBy string    `db:"updated_by"` // 更新人
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func NewSysRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleMenuModel {
	return &defaultSysRoleMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_role_menu`",
	}
}

func (m *defaultSysRoleMenuModel) Insert(data *SysRoleMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysRoleMenuRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.RoleId, data.MenuId, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt)

	return ret, err
}

func (m *defaultSysRoleMenuModel) FindOne(id int64) (*SysRoleMenu, error) {
	sysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheSysRoleMenuIdPrefix, id)
	var resp SysRoleMenu
	err := m.QueryRow(&resp, sysRoleMenuIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
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

func (m *defaultSysRoleMenuModel) Update(data *SysRoleMenu) error {
	sysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheSysRoleMenuIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleMenuRowsWithPlaceHolder)
		return conn.Exec(query, data.RoleId, data.MenuId, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.Id)
	}, sysRoleMenuIdKey)
	return err
}

func (m *defaultSysRoleMenuModel) Delete(id int64) error {

	sysRoleMenuIdKey := fmt.Sprintf("%s%v", cacheSysRoleMenuIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysRoleMenuIdKey)
	return err
}

func (m *defaultSysRoleMenuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysRoleMenuIdPrefix, primary)
}

func (m *defaultSysRoleMenuModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
	return conn.QueryRow(v, query, primary)
}
