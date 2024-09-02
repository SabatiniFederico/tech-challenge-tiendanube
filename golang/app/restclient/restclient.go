package restclient

import (
	"fmt"
	"github.com/fsabatini/tech-challenge-tiendanube/app/config"
	"github.com/fsabatini/tech-challenge-tiendanube/app/domain"
	"github.com/go-resty/resty/v2"
	"log"
)

type (
	TransactionsRestClient struct {
		RestClient *resty.Client
	}
)

func (transactionRest TransactionsRestClient) GetTransactions() []domain.Transaction {
	return nil
}

func (transactionRest TransactionsRestClient) PostTransaction(transaction domain.Transaction) error {
	result, err := transactionRest.RestClient.R().SetBody(transaction).Post(config.URL + config.PATHTRANSACTIONS)
	log.Println(fmt.Sprint("results are: %v", result))
	return err
}
func (transactionRest TransactionsRestClient) PostPayable(payable domain.Payable) error {
	_, err := transactionRest.RestClient.R().SetBody(payable).Post(config.URL + config.PATHPAYABLES)
	return err
}
