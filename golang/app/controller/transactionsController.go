package controller

import (
	"github.com/fsabatini/tech-challenge-tiendanube/app/domain"
	"github.com/gin-gonic/gin"
)

type (
	transactionsService interface {
		CalculateTransactionAndPayable(transaction domain.Transaction) error
	}

	TransactionsController struct {
		TransactionsService transactionsService
	}
)

func (tc TransactionsController) CalculatePayment(c *gin.Context) {

	entity := &domain.Transaction{}
	err := c.BindJSON(entity)

	if err != nil {
		c.JSON(400, gin.H{"reason": "bad format of transaction"})
		return
	}
	//Validations
	errService := tc.TransactionsService.CalculateTransactionAndPayable(*entity)

	if errService != nil {
		c.JSON(500, gin.H{"reason": "Error on transactions service"})
		return
	}

	c.JSON(200, gin.H{"calculation": "done"})
}

func (tc TransactionsController) ListTransactions(c *gin.Context) {

	c.JSON(200, gin.H{"calculation": "done"})
}
