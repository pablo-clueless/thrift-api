package router

import (
	"github.com/gin-gonic/gin"
	"thrift.com/m/services"
)

func AuthRouter(router *gin.RouterGroup) *gin.RouterGroup {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", services.Signup())
		auth.POST("/signin", services.Signin())
	}
	return auth
}
