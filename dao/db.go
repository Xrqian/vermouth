package dao

import (
	"time"
	"vermouth/conf"
	"vermouth/model"
	"vermouth/util"
)

func Init() {
	// Init Mysql and Migrate the schema
	mysqlConnect := util.GetConnect(
		conf.Cfg.MysqlUserName,
		conf.Cfg.MysqlDbPass,
		conf.Cfg.MysqlDbHost,
		conf.Cfg.MysqlDbPort,
		conf.Cfg.MysqlDbName,
	)
	mysqlConnect.DB.DB().SetMaxOpenConns(conf.Cfg.MysqlMaxOpenConns)
	mysqlConnect.DB.DB().SetConnMaxLifetime(10 * time.Minute)
	mysqlConnect.DB.AutoMigrate(&model.User{}, &model.PayRecord{})
}
