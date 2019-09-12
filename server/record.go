package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vermouth/dao"
	"vermouth/model"
	"vermouth/util"
)

func listRecord(c *gin.Context) {
	handle := dao.PayRecordDaoImpl{}
	page := util.ParseInt(c.DefaultQuery("page", "1"))
	pageSize := util.ParseInt(c.DefaultQuery("size", "10"))
	parameter := model.Parameter{
		Page:     page,
		PageSize: pageSize,
	}
	recordList, err := handle.List(&parameter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recordList)
}

func retrieveRecord(c *gin.Context) {
	handle := dao.PayRecordDaoImpl{}
	id := util.ParseInt(c.Param("id"))
	record, err := handle.Retrieve(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
}
