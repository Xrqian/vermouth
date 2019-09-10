package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vermouth/dao"
	"vermouth/model"
	"strconv"
)

func parseInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}

func listUser(c *gin.Context) {
	handle := dao.UserDaoImpl{}
	page := parseInt(c.DefaultQuery("page", "1"))
	pageSize := parseInt(c.DefaultQuery("size", "10"))
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

func listUserRecord(c *gin.Context) {
	c.String(http.StatusOK, "listUserRecord")
}
