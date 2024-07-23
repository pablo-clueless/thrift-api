package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"thrift.com/m/common"
	"thrift.com/m/database"
	"thrift.com/m/router"
)

func Setup() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := database.Connect()
	var err error
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			common.Logger("ERROR", "Error disconnecting from MongoDB: %v", err)
			return
		}
	}()

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	port := ":8080"
	api := app.Group("/thrift/v1")

	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Thrift Personal Finance Tracker!",
		})
	})

	router.AuthRouter(api)
	router.AccountRouter(api)
	router.TransactionRouter(api)

	err = app.Run(port)
	if err != nil {
		common.Logger("ERROR", "Error running the application:", err)
		return err
	}
	common.Logger("INFO", "Application running on port:", nil)
	return nil
}
