package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	cartsFieldNames          = builder.RawFieldNames(&Carts{})
	cartsRows                = strings.Join(cartsFieldNames, ",")
	cartsRowsExpectAutoSet   = strings.Join(stringx.Remove(cartsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	cartsRowsWithPlaceHolder = strings.Join(stringx.Remove(cartsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheCartsIdPrefix = "cache:carts:id:"
)

type (
	CartsModel interface {
		Insert(data *Carts) (sql.Result, error)
		FindOne(id int64) (*Carts, error)
		Update(data *Carts) error
		Delete(id int64) error
		FindAllByUid(uid int64) ([]*Carts, error)
	}

	defaultCartsModel struct {
		sqlc.CachedConn
		table string
	}

	Carts struct {
		Id         int64     `db:"id"`
		Uid        int64     `db:"uid"`  // uid
		Gid        int64     `db:"gid"`  // 商品id
		Name       string    `db:"name"` // 产品名称
		Num        int64     `db:"num"`  // 数量
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewCartsModel(conn sqlx.SqlConn, c cache.CacheConf) CartsModel {
	return &defaultCartsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`carts`",
	}
}

func (m *defaultCartsModel) Insert(data *Carts) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, cartsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Uid, data.Gid, data.Name, data.Num)

	return ret, err
}

func (m *defaultCartsModel) FindOne(id int64) (*Carts, error) {
	cartsIdKey := fmt.Sprintf("%s%v", cacheCartsIdPrefix, id)
	var resp Carts
	err := m.QueryRow(&resp, cartsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cartsRows, m.table)
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

func (m *defaultCartsModel) Update(data *Carts) error {
	cartsIdKey := fmt.Sprintf("%s%v", cacheCartsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cartsRowsWithPlaceHolder)
		return conn.Exec(query, data.Uid, data.Gid, data.Name, data.Num, data.Id)
	}, cartsIdKey)
	return err
}

func (m *defaultCartsModel) Delete(id int64) error {

	cartsIdKey := fmt.Sprintf("%s%v", cacheCartsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, cartsIdKey)
	return err
}

func (m *defaultCartsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCartsIdPrefix, primary)
}

func (m *defaultCartsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cartsRows, m.table)
	return conn.QueryRow(v, query, primary)
}

func (m *defaultCartsModel) FindAllByUid(uid int64) ([]*Carts, error) {
	var resp []*Carts

	query := fmt.Sprintf("select %s from %s where `uid` = ?", cartsRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
