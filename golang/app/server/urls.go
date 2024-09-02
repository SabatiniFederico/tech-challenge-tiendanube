package server

import (
	"github.com/gin-gonic/gin"
)

func mapUrlsToControllers(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

}
