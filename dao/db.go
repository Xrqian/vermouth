package dao

import (
	"time"
	"vermouth/conf"
	"vermouth/model"
	"vermouth/util"
)

func Init() {
	// Init Mysql and Migrate the schema
	config := conf.GetCoreConfig()
	mysqlClient := util.GetConnect(config.MysqlUserName, config.MysqlDbPass, config.MysqlDbHost, config.MysqlDbPort, config.MysqlDbName)
	mysqlClient.Db.DB().SetMaxOpenConns(config.MysqlMaxOpenConns)
	mysqlClient.Db.DB().SetConnMaxLifetime(10 * time.Minute)
	mysqlClient.Db.AutoMigrate(&model.User{}, &model.PayRecord{})
}
