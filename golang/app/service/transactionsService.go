package service

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/domain"
	"log"
)

type (
	transactionsAndPayablesRestClient interface {
		GetTransactions() []domain.Transaction
		PostTransaction(transaction domain.Transaction) error
		PostPayable(payable domain.Payable) error
	}

	TransactionsService struct {
		Restclient transactionsAndPayablesRestClient
	}
)

func (transactionService TransactionsService) CalculateTransactionAndPayable(transaction domain.Transaction) error {
	log.Println("Executing service transaction")
	err := transactionService.Restclient.PostTransaction(transaction)
	return err
}
