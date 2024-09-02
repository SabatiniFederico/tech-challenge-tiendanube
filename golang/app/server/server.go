package server

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	bootstrap(router)

	return router
}
