package service

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/service"
	"github.com/go-resty/resty/v2"
)

func testTransaction() {
	client := resty.New()
	service := &service.TransactionsService{
		Restclient: client}
}
