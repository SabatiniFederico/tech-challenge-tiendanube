package server

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func bootstrap(router *gin.Engine) {

	transactionsController := &controller.TransactionsController{}

	log.Println("Initializing controllers")
	mapUrlsToControllers(router, transactionsController)

}
