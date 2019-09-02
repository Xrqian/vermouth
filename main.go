package main

import (
	"github.com/blazehu/vermouth/util"
	"github.com/blazehu/vermouth/model"
	"github.com/blazehu/vermouth/conf"
)

func main() {
	config := conf.GetCoreConfig()
	mysqlClient := util.GetConnect(config.MysqlUserName, config.MysqlDbPass, config.MysqlDbHost, config.MysqlDbPort, config.MysqlDbName)
	// Migrate the schema
	mysqlClient.Db.AutoMigrate(&model.User{})
	defer mysqlClient.Db.Close()
}
