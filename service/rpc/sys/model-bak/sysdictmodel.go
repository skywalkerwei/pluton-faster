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
	sysDictFieldNames          = builder.RawFieldNames(&SysDict{})
	sysDictRows                = strings.Join(sysDictFieldNames, ",")
	sysDictRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysDictRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysDictIdPrefix = "cache:sysDict:id:"
)

type (
	SysDictModel interface {
		Insert(data *SysDict) (sql.Result, error)
		FindOne(id int64) (*SysDict, error)
		Update(data *SysDict) error
		Delete(id int64) error
	}

	defaultSysDictModel struct {
		sqlc.CachedConn
		table string
	}

	SysDict struct {
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

func NewSysDictModel(conn sqlx.SqlConn, c cache.CacheConf) SysDictModel {
	return &defaultSysDictModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_dict`",
	}
}

func (m *defaultSysDictModel) Insert(data *SysDict) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysDictRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Value, data.Label, data.Tp, data.Description, data.Sort, data.Remarks, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultSysDictModel) FindOne(id int64) (*SysDict, error) {
	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, id)
	var resp SysDict
	err := m.QueryRow(&resp, sysDictIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
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

func (m *defaultSysDictModel) Update(data *SysDict) error {
	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDictRowsWithPlaceHolder)
		return conn.Exec(query, data.Value, data.Label, data.Tp, data.Description, data.Sort, data.Remarks, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, sysDictIdKey)
	return err
}

func (m *defaultSysDictModel) Delete(id int64) error {

	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysDictIdKey)
	return err
}

func (m *defaultSysDictModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysDictIdPrefix, primary)
}

func (m *defaultSysDictModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
	return conn.QueryRow(v, query, primary)
}
