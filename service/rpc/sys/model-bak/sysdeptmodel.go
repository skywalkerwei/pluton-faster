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
	sysDeptFieldNames          = builder.RawFieldNames(&SysDept{})
	sysDeptRows                = strings.Join(sysDeptFieldNames, ",")
	sysDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDeptFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDeptFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysDeptIdPrefix = "cache:sysDept:id:"
)

type (
	SysDeptModel interface {
		Insert(data *SysDept) (sql.Result, error)
		FindOne(id int64) (*SysDept, error)
		Update(data *SysDept) error
		Delete(id int64) error
	}

	defaultSysDeptModel struct {
		sqlc.CachedConn
		table string
	}

	SysDept struct {
		Id        int64        `db:"id"`         // 编号
		Name      string       `db:"name"`       // 机构名称
		ParentId  int64        `db:"parent_id"`  // 上级机构ID，一级机构为0
		OrderNum  int64        `db:"order_num"`  // 排序
		CreatedBy string       `db:"created_by"` // 创建人
		UpdatedBy string       `db:"updated_by"` // 更新人
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt time.Time    `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func NewSysDeptModel(conn sqlx.SqlConn, c cache.CacheConf) SysDeptModel {
	return &defaultSysDeptModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_dept`",
	}
}

func (m *defaultSysDeptModel) Insert(data *SysDept) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysDeptRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.ParentId, data.OrderNum, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultSysDeptModel) FindOne(id int64) (*SysDept, error) {
	sysDeptIdKey := fmt.Sprintf("%s%v", cacheSysDeptIdPrefix, id)
	var resp SysDept
	err := m.QueryRow(&resp, sysDeptIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDeptRows, m.table)
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

func (m *defaultSysDeptModel) Update(data *SysDept) error {
	sysDeptIdKey := fmt.Sprintf("%s%v", cacheSysDeptIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDeptRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.ParentId, data.OrderNum, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, sysDeptIdKey)
	return err
}

func (m *defaultSysDeptModel) Delete(id int64) error {

	sysDeptIdKey := fmt.Sprintf("%s%v", cacheSysDeptIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysDeptIdKey)
	return err
}

func (m *defaultSysDeptModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysDeptIdPrefix, primary)
}

func (m *defaultSysDeptModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDeptRows, m.table)
	return conn.QueryRow(v, query, primary)
}
