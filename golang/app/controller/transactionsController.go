package controller

import "github.com/gin-gonic/gin"

type (
	TransactionsController struct {
	}
)

func (tc TransactionsController) CalculatePayment(c *gin.Context) {
	c.JSON(200, gin.H{"calculation": "done"})
}
