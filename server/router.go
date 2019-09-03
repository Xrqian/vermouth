package server

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", listUser)
	v1.GET("/record", listUserRecord)
}
