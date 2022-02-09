package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

//Dbc ，
func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "User:Password@tcp(Host:Port)/Db?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil
	}
	//sqlDB, _ := db.DB()
	//// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	//sqlDB.SetMaxIdleConns(10)
	////// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDB.SetMaxOpenConns(100)
	////// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDB.SetConnMaxLifetime(time.Hour)
	readDbIsOpen := 1

	if readDbIsOpen == 1 {
		if val, err := getDbDialector(sqlType, "Read", dbConf...); err == nil {
			dbDialector = val
		}
		resolverConf := dbresolver.Config{
			Replicas: []gorm.Dialector{dbDialector}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},     // sources/replicas 负载均衡策略适用于
		}
		err = gormDb.Use(dbresolver.Register(resolverConf).SetConnMaxIdleTime(time.Second * 30).
			SetConnMaxLifetime(variable.ConfigGormv2Yml.GetDuration("Gormv2."+sqlType+".Read.SetConnMaxLifetime") * time.Second).
			SetMaxIdleConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Read.SetMaxIdleConns")).
			SetMaxOpenConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Read.SetMaxOpenConns")))
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
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(variable.ConfigGormv2Yml.GetDuration("Gormv2."+sqlType+".Write.SetConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Write.SetMaxIdleConns"))
		rawDb.SetMaxOpenConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Write.SetMaxOpenConns"))
		return gormDb, nil
	}

	return db
}
