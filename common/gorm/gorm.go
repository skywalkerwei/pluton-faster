package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func GetSqlDriver(conf Conf) *gorm.DB {
	var dbDialector gorm.Dialector
	if val, err := getDbDialector("w", conf); err == nil {
		dbDialector = val
	}
	gormDb, err := gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil
	}

	if conf.IsOpenReadDb {

		if val, err := getDbDialector("r", conf); err == nil {
			dbDialector = val
		}
		resolverConf := dbresolver.Config{
			Replicas: []gorm.Dialector{dbDialector}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},     // sources/replicas 负载均衡策略适用于
		}
		err = gormDb.Use(
			dbresolver.Register(resolverConf).
				SetConnMaxLifetime(60 * time.Second).
				SetMaxIdleConns(conf.Read.SetMaxIdleConns).
				SetMaxOpenConns(conf.Read.SetMaxOpenConns),
		)
		if err != nil {
			return nil
		}
	}

	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = gormDb.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError)

	// https://github.com/go-gorm/gorm/issues/4838
	_ = gormDb.Callback().Create().Before("gorm:before_create").Register("CreateBeforeHook", CreateBeforeHook)
	// 为了完美支持gorm的一系列回调函数
	_ = gormDb.Callback().Update().Before("gorm:before_update").Register("UpdateBeforeHook", UpdateBeforeHook)

	// 为主连接设置连接池(43行返回的数据库驱动指针)
	if rawDb, err := gormDb.DB(); err == nil {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(60 * time.Second)
		rawDb.SetMaxIdleConns(conf.Write.SetMaxIdleConns)
		rawDb.SetMaxOpenConns(conf.Write.SetMaxOpenConns)
		return gormDb
	}

	return gormDb
}

// 获取一个数据库方言(Dialector),通俗的说就是根据不同的连接参数，获取具体的一类数据库的连接指针
func getDbDialector(rw string, conf Conf) (gorm.Dialector, error) {
	var dbDialector gorm.Dialector
	dsn := getDsn(rw, conf)
	dbDialector = mysql.Open(dsn)
	return dbDialector, nil
}

func getDsn(rw string, conf Conf) string {
	if rw == "r" {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",
			conf.Read.User,
			conf.Read.Pass,
			conf.Read.Host,
			conf.Read.Port,
			conf.Read.DataBase,
			conf.Read.Charset)
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",
			conf.Write.User,
			conf.Write.Pass,
			conf.Write.Host,
			conf.Write.Port,
			conf.Write.DataBase,
			conf.Write.Charset)
	}
}
