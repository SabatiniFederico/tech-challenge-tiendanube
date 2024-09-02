package server

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/controller"
	"github.com/fsabatini/tech-challenge-tiendanube/app/restclient"
	"github.com/fsabatini/tech-challenge-tiendanube/app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
)

func bootstrap(router *gin.Engine) {

	client := resty.New()
	transactionsController := &controller.TransactionsController{
		TransactionsService: &service.TransactionsService{
			Restclient: &restclient.TransactionsRestClient{
				RestClient: client,
			},
		},
	}

	log.Println("Initializing controllers")
	mapUrlsToControllers(router, transactionsController)

}
