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
	sysJobFieldNames          = builder.RawFieldNames(&SysJob{})
	sysJobRows                = strings.Join(sysJobFieldNames, ",")
	sysJobRowsExpectAutoSet   = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysJobRowsWithPlaceHolder = strings.Join(stringx.Remove(sysJobFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysJobIdPrefix = "cache:sysJob:id:"
)

type (
	SysJobModel interface {
		Insert(data *SysJob) (sql.Result, error)
		FindOne(id int64) (*SysJob, error)
		Update(data *SysJob) error
		Delete(id int64) error
	}

	defaultSysJobModel struct {
		sqlc.CachedConn
		table string
	}

	SysJob struct {
		Id        int64         `db:"id"`         // 编号
		JobName   string        `db:"job_name"`   // 职位名称
		OrderNum  int64         `db:"order_num"`  // 排序
		Remarks   string        `db:"remarks"`    // 备注
		CreatedBy string        `db:"created_by"` // 创建人
		UpdatedBy string        `db:"updated_by"` // 更新人
		CreatedAt time.Time     `db:"created_at"`
		UpdatedAt time.Time     `db:"updated_at"`
		DeletedAt sql.NullInt64 `db:"deleted_at"` // 删除时间
	}
)

func NewSysJobModel(conn sqlx.SqlConn, c cache.CacheConf) SysJobModel {
	return &defaultSysJobModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_job`",
	}
}

func (m *defaultSysJobModel) Insert(data *SysJob) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysJobRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.JobName, data.OrderNum, data.Remarks, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultSysJobModel) FindOne(id int64) (*SysJob, error) {
	sysJobIdKey := fmt.Sprintf("%s%v", cacheSysJobIdPrefix, id)
	var resp SysJob
	err := m.QueryRow(&resp, sysJobIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysJobRows, m.table)
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

func (m *defaultSysJobModel) Update(data *SysJob) error {
	sysJobIdKey := fmt.Sprintf("%s%v", cacheSysJobIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysJobRowsWithPlaceHolder)
		return conn.Exec(query, data.JobName, data.OrderNum, data.Remarks, data.CreatedBy, data.UpdatedBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, sysJobIdKey)
	return err
}

func (m *defaultSysJobModel) Delete(id int64) error {

	sysJobIdKey := fmt.Sprintf("%s%v", cacheSysJobIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysJobIdKey)
	return err
}

func (m *defaultSysJobModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysJobIdPrefix, primary)
}

func (m *defaultSysJobModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysJobRows, m.table)
	return conn.QueryRow(v, query, primary)
}
