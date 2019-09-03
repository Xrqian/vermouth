package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vermouth/conf"
	"vermouth/model"
	"vermouth/util"
)

func listUser(c *gin.Context) {
	// Init Mysql and Migrate the schema
	mysqlConnect := util.GetConnect(
		conf.Cfg.MysqlUserName,
		conf.Cfg.MysqlDbPass,
		conf.Cfg.MysqlDbHost,
		conf.Cfg.MysqlDbPort,
		conf.Cfg.MysqlDbName,
	)
	var users = make([]model.User, 0)
	mysqlConnect.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func listUserRecord(c *gin.Context) {
	c.String(http.StatusOK, "listUserRecord")
}
