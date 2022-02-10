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
	sysMenuFieldNames          = builder.RawFieldNames(&SysMenu{})
	sysMenuRows                = strings.Join(sysMenuFieldNames, ",")
	sysMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysMenuIdPrefix = "cache:sysMenu:id:"
)

type (
	SysMenuModel interface {
		Insert(data *SysMenu) (sql.Result, error)
		FindOne(id int64) (*SysMenu, error)
		Update(data *SysMenu) error
		Delete(id int64) error
	}

	defaultSysMenuModel struct {
		sqlc.CachedConn
		table string
	}

	SysMenu struct {
		Id            int64        `db:"id"`        // 编号
		Name          string       `db:"name"`      // 菜单名称
		ParentId      int64        `db:"parent_id"` // 父菜单ID，一级菜单为0
		Url           string       `db:"url"`
		Perms         string       `db:"perms"`          // 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
		Tp            int64        `db:"tp"`             // 类型   0：目录   1：菜单   2：按钮
		Icon          string       `db:"icon"`           // 菜单图标
		OrderNum      int64        `db:"order_num"`      // 排序
		VuePath       string       `db:"vue_path"`       // vue系统的path
		VueComponent  string       `db:"vue_component"`  // vue的页面
		VueIcon       string       `db:"vue_icon"`       // vue的图标
		VueRedirect   string       `db:"vue_redirect"`   // vue的路由重定向
		BackgroundUrl string       `db:"background_url"` // 后台地址
		CreatedBy     string       `db:"created_by"`     // 创建人
		UpdateBy      string       `db:"update_by"`      // 更新人
		CreatedAt     time.Time    `db:"created_at"`
		UpdatedAt     time.Time    `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func NewSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf) SysMenuModel {
	return &defaultSysMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_menu`",
	}
}

func (m *defaultSysMenuModel) Insert(data *SysMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysMenuRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.ParentId, data.Url, data.Perms, data.Tp, data.Icon, data.OrderNum, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect, data.BackgroundUrl, data.CreatedBy, data.UpdateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultSysMenuModel) FindOne(id int64) (*SysMenu, error) {
	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, id)
	var resp SysMenu
	err := m.QueryRow(&resp, sysMenuIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
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

func (m *defaultSysMenuModel) Update(data *SysMenu) error {
	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMenuRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.ParentId, data.Url, data.Perms, data.Tp, data.Icon, data.OrderNum, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect, data.BackgroundUrl, data.CreatedBy, data.UpdateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, sysMenuIdKey)
	return err
}

func (m *defaultSysMenuModel) Delete(id int64) error {

	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysMenuIdKey)
	return err
}

func (m *defaultSysMenuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, primary)
}

func (m *defaultSysMenuModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
	return conn.QueryRow(v, query, primary)
}
