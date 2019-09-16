package server

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", listUser)
	v1.GET("/user/:id", retrieveUser)
	v1.POST("/user", createUser)
	v1.GET("/record", listRecord)
	v1.GET("/record/:id", retrieveRecord)
	v1.POST("/record", createRecord)
}
