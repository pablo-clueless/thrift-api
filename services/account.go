package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"thrift.com/m/common"
	"thrift.com/m/database"
	"thrift.com/m/models"
)

var accounts *mongo.Collection = database.GetCollection(database.Database, "accounts")

func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var account models.AccountProps
		var user models.UserProps
		defer cancel()

		if err := c.BindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid request payload!",
				Data:    nil,
			})
			return
		}

		userId, _ := primitive.ObjectIDFromHex(account.UserId)
		err := users.FindOne(ctx, bson.M{"id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, common.Response{
				Error:   true,
				Code:    404,
				Message: "User not found!",
				Data:    nil,
			})
			return
		}

		newAccount := models.AccountProps{
			Id:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			Balance:   0,
			Name:      account.Name,
			UserId:    account.UserId,
		}

		_, err = accounts.InsertOne(ctx, newAccount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Something went wrong!",
				Data:    nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Error:   false,
			Code:    201,
			Message: "Account created successfully!",
			Data:    nil,
		})
	}
}

func GetAccountsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var accountsList []models.AccountProps
		defer cancel()

		id := c.Param("id")
		cursor, err := accounts.Find(ctx, bson.M{"user_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Something went wrong!",
				Data:    nil,
			})
			return
		}

		if err = cursor.All(ctx, &accountsList); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Something went wrong!",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Error:   false,
			Code:    200,
			Message: "Accounts found!",
			Data:    accountsList,
		})
	}
}

func GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var account models.AccountProps
		defer cancel()

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := accounts.FindOne(ctx, bson.M{"_id": id}).Decode(&account)
		if err != nil {
			c.JSON(http.StatusNotFound, common.Response{
				Error:   true,
				Code:    404,
				Message: "Account not found!",
				Data:    nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Error:   false,
			Code:    200,
			Message: "Account found!",
			Data:    account,
		})
	}
}

func UpdateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var updates models.AccountProps
		defer cancel()

		if err := c.BindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid request payload!",
				Data:    nil,
			})
			return
		}

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var account models.AccountProps
		err := accounts.FindOne(ctx, bson.M{"_id": id}).Decode(&account)
		if err != nil {
			c.JSON(http.StatusNotFound, common.Response{
				Error:   true,
				Code:    404,
				Message: "Account not found!",
				Data:    nil,
			})
			return
		}

		var updatedName string
		if updates.Name != "" {
			updatedName = updates.Name
		} else {
			updatedName = account.Name
		}
		update := bson.M{"$set": bson.M{
			"name":       updatedName,
			"balance":    account.Balance + updates.Balance,
			"updated_at": time.Now(),
		}}

		updated, err := accounts.UpdateOne(ctx, bson.M{"_id": id}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Something went wrong!",
				Data:    nil,
			})
			return
		}
		if updated.MatchedCount == 0 {
			err := accounts.FindOne(ctx, id).Decode(&account)
			if err != nil {
				c.JSON(http.StatusNotFound, common.Response{
					Error:   true,
					Code:    404,
					Message: "Unable to update account!",
					Data:    nil,
				})
				return
			}
		}

		c.JSON(http.StatusOK, common.Response{
			Error:   false,
			Code:    200,
			Message: "Account updated successfully!",
			Data:    nil,
		})
	}
}

func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var account models.AccountProps
		defer cancel()

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := accounts.FindOne(ctx, bson.M{"_id": id}).Decode(&account)
		if err != nil {
			c.JSON(http.StatusNotFound, common.Response{
				Error:   true,
				Code:    404,
				Message: "Account not found!",
				Data:    nil,
			})
			return
		}
		_, err = accounts.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Something went wrong!",
				Data:    nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Error:   false,
			Code:    200,
			Message: "Account deleted successfully!",
			Data:    nil,
		})
	}
}
