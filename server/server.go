package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vermouth/conf"
)

func Init() *http.Server {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	InitRouter(router)

	addr := fmt.Sprintf(":%s", conf.Cfg.HttpPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return srv
}
