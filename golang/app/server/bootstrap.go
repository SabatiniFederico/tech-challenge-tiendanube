package server

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/controller"
	"github.com/fsabatini/tech-challenge-tiendanube/app/service"
	"github.com/gin-gonic/gin"
	"log"
)

func bootstrap(router *gin.Engine) {

	transactionsController := &controller.TransactionsController{
		TransactionsService: &service.TransactionsService{},
	}

	log.Println("Initializing controllers")
	mapUrlsToControllers(router, transactionsController)

}
