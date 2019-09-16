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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userList)
}

func retrieveUser(c *gin.Context) {
	handle := dao.UserDaoImpl{}

	id := util.ParseInt(c.Param("id"))
	user, err := handle.Retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	handle := dao.UserDaoImpl{}

	userData := model.User{}
	err := c.ShouldBind(&userData)

	user, err := handle.Create(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
