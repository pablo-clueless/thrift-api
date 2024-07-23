package router

import (
	"github.com/gin-gonic/gin"
	"thrift.com/m/middleware"
	"thrift.com/m/services"
)

func AccountRouter(router *gin.RouterGroup) *gin.RouterGroup {
	account := router.Group("/account")
	{
		account.POST("", middleware.Authorize(), services.CreateAccount())
		account.GET("/user/:id", middleware.Authorize(), services.GetAccountsByUserId())
		account.GET("/:id", middleware.Authorize(), services.GetAccount())
		account.PUT("/:id", middleware.Authorize(), services.UpdateAccount())
		account.DELETE("/:id", middleware.Authorize(), services.DeleteAccount())
	}
	return account
}
