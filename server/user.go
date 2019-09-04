package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vermouth/dao"
	"vermouth/model"
)

func listUser(c *gin.Context) {
	db := dao.MysqlDB
	users := make([]model.User, 0)
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func listUserRecord(c *gin.Context) {
	c.String(http.StatusOK, "listUserRecord")
}
