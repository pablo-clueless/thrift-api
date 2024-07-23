package router

import (
	"github.com/gin-gonic/gin"
	"thrift.com/m/middleware"
	"thrift.com/m/services"
)

func TransactionRouter(router *gin.RouterGroup) gin.RouterGroup {
	transaction := router.Group("/transactions")
	{
		transaction.POST("", middleware.Authorize(), services.CreateTransaction())
		transaction.GET("/account/:id", middleware.Authorize(), services.GetTransactionsByAccountId())
		transaction.GET("/:id", middleware.Authorize(), services.GetTransactionById())
		transaction.DELETE("/:id", middleware.Authorize(), services.DeleteTransaction())
	}
	return *transaction
}
