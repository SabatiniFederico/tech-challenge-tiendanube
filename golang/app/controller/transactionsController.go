package controller

import "github.com/gin-gonic/gin"

type (
	transactionsService interface {
		CalculateTransactionAndPayable()
	}

	TransactionsController struct {
		TransactionsService transactionsService
	}
)

func (tc TransactionsController) CalculatePayment(c *gin.Context) {

	tc.TransactionsService.CalculateTransactionAndPayable()
	c.JSON(200, gin.H{"calculation": "done"})
}
