package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vermouth/dao"
	"vermouth/model"
	"vermouth/util"
)

func listUser(c *gin.Context) {
	handle := dao.UserDaoImpl{}
	page := util.ParseInt(c.DefaultQuery("page", "1"))
	pageSize := util.ParseInt(c.DefaultQuery("size", "10"))
	parameter := model.Parameter{
		Page:     page,
		PageSize: pageSize,
	}
	userList, err := handle.List(&parameter)
	if err != nil {
		c.JSON(http.StatusBadRequest, userList)
	}
	c.JSON(http.StatusOK, userList)
}

func retrieveUser(c *gin.Context) {
	handle := dao.UserDaoImpl{}
	id := util.ParseInt(c.Param("id"))
	user, err := handle.Retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, user)
	}
	c.JSON(http.StatusOK, user)
}

func listUserRecord(c *gin.Context) {
	handle := dao.PayRecordDaoImpl{}
	page := util.ParseInt(c.DefaultQuery("page", "1"))
	pageSize := util.ParseInt(c.DefaultQuery("size", "10"))
	parameter := model.Parameter{
		Page:     page,
		PageSize: pageSize,
	}
	recordList, err := handle.List(&parameter)
	if err != nil {
		c.JSON(http.StatusBadRequest, recordList)
	}
	c.JSON(http.StatusOK, recordList)
}
