package service

type (
	transactionsAndPayablesRestClient interface {
		GetTransactions()
		PostTransaction()
		PostPayable()
	}

	TransactionsService struct {
		Restclient transactionsAndPayablesRestClient
	}
)

func (transactionService TransactionsService) CalculateTransactionAndPayable() {

}
