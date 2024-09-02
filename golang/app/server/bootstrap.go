package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func bootstrap(router *gin.Engine) {

	log.Println("Initializing controllers")
	mapUrlsToControllers(router)

}
