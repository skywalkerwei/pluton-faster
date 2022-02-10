package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.SysUserModel
	UserRoleModel model.SysUserRoleModel
	RoleModel     model.SysRoleModel
	RoleMenuModel model.SysRoleMenuModel
	MenuModel     model.SysMenuModel
	DictModel     model.SysDictModel
	DeptModel     model.SysDeptModel
	LoginLogModel model.SysLoginLogModel
	SysLogModel   model.SysLogModel
	ConfigModel   model.SysConfigModel
	JobModel      model.SysJobModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UserModel:     model.NewSysUserModel(conn, c.CacheRedis),
		UserRoleModel: model.NewSysUserRoleModel(conn, c.CacheRedis),
		RoleModel:     model.NewSysRoleModel(conn, c.CacheRedis),
		RoleMenuModel: model.NewSysRoleMenuModel(conn, c.CacheRedis),
		MenuModel:     model.NewSysMenuModel(conn, c.CacheRedis),
		DictModel:     model.NewSysDictModel(conn, c.CacheRedis),
		DeptModel:     model.NewSysDeptModel(conn, c.CacheRedis),
		LoginLogModel: model.NewSysLoginLogModel(conn, c.CacheRedis),
		SysLogModel:   model.NewSysLogModel(conn, c.CacheRedis),
		ConfigModel:   model.NewSysConfigModel(conn, c.CacheRedis),
		JobModel:      model.NewSysJobModel(conn, c.CacheRedis),
	}
}
