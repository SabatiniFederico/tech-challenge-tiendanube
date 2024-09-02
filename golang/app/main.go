package main

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/server"
	"log"
)

func main() {
	err := server.NewServer().Run(":8080")
	if err != nil {
		log.Fatal("Error on launching server")
	}
}
