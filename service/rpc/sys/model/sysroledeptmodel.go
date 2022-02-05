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
	sysRoleDeptFieldNames          = builder.RawFieldNames(&SysRoleDept{})
	sysRoleDeptRows                = strings.Join(sysRoleDeptFieldNames, ",")
	sysRoleDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleDeptFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleDeptFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysRoleDeptIdPrefix = "cache:sysRoleDept:id:"
)

type (
	SysRoleDeptModel interface {
		Insert(data *SysRoleDept) (sql.Result, error)
		FindOne(id int64) (*SysRoleDept, error)
		Update(data *SysRoleDept) error
		Delete(id int64) error
	}

	defaultSysRoleDeptModel struct {
		sqlc.CachedConn
		table string
	}

	SysRoleDept struct {
		Id        int64     `db:"id"`         // 编号
		RoleId    int64     `db:"role_id"`    // 角色ID
		DeptId    int64     `db:"dept_id"`    // 机构ID
		CreatedBy string    `db:"created_by"` // 创建人
		UpdatedBy string    `db:"updated_by"` // 更新人
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func NewSysRoleDeptModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleDeptModel {
	return &defaultSysRoleDeptModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_role_dept`",
	}
}

func (m *defaultSysRoleDeptModel) Insert(data *SysRoleDept) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysRoleDeptRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.RoleId, data.DeptId, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt)

	return ret, err
}

func (m *defaultSysRoleDeptModel) FindOne(id int64) (*SysRoleDept, error) {
	sysRoleDeptIdKey := fmt.Sprintf("%s%v", cacheSysRoleDeptIdPrefix, id)
	var resp SysRoleDept
	err := m.QueryRow(&resp, sysRoleDeptIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleDeptRows, m.table)
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

func (m *defaultSysRoleDeptModel) Update(data *SysRoleDept) error {
	sysRoleDeptIdKey := fmt.Sprintf("%s%v", cacheSysRoleDeptIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleDeptRowsWithPlaceHolder)
		return conn.Exec(query, data.RoleId, data.DeptId, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.Id)
	}, sysRoleDeptIdKey)
	return err
}

func (m *defaultSysRoleDeptModel) Delete(id int64) error {

	sysRoleDeptIdKey := fmt.Sprintf("%s%v", cacheSysRoleDeptIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysRoleDeptIdKey)
	return err
}

func (m *defaultSysRoleDeptModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysRoleDeptIdPrefix, primary)
}

func (m *defaultSysRoleDeptModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleDeptRows, m.table)
	return conn.QueryRow(v, query, primary)
}
