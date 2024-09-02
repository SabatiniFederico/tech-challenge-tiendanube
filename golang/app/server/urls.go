package server

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/controller"
	"github.com/gin-gonic/gin"
)

func mapUrlsToControllers(router *gin.Engine, transactionsController *controller.TransactionsController) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/payables", transactionsController.CalculatePayment)
}
