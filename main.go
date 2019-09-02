package main

import (
	"vermouth/util"
	"vermouth/model"
	"vermouth/conf"
)

func main() {
	config := conf.GetCoreConfig()
	mysqlClient := util.GetConnect(config.MysqlUserName, config.MysqlDbPass, config.MysqlDbHost, config.MysqlDbPort, config.MysqlDbName)
	// Migrate the schema
	mysqlClient.Db.AutoMigrate(&model.User{})
	defer mysqlClient.Db.Close()
}
