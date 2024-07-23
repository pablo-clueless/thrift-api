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

var transactions *mongo.Collection = database.GetCollection(database.Database, "transactions")

func CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction *models.TransactionProps
		var account *models.AccountProps
		defer cancel()

		if err := c.BindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid request payload!",
				Data:    nil,
			})
			return
		}

		id, _ := primitive.ObjectIDFromHex(transaction.AccountId)
		err := accounts.FindOne(ctx, bson.M{"_id": id}).Decode(&account)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    404,
				Message: "Account not found!",
				Data:    nil,
			})
			return
		}

		newTransaction := models.TransactionProps{
			Id:          primitive.NewObjectID(),
			CreatedAt:   time.Now(),
			AccountId:   transaction.AccountId,
			Amount:      transaction.Amount,
			Category:    transaction.Category,
			Date:        time.Now(),
			Description: transaction.Description,
			Status:      "done",
			Type:        transaction.Type,
		}
		_, err = transactions.InsertOne(ctx, newTransaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Error creating transaction",
				Data:    nil,
			})
			return
		}

		var update bson.M
		if transaction.Type == "incoming" {
			update = bson.M{"$set": bson.M{
				"name":    account.Name,
				"balance": account.Balance + transaction.Amount,
			}}
		} else {
			update = bson.M{"$set": bson.M{
				"name":    account.Name,
				"balance": account.Balance + transaction.Amount,
			}}
		}
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

		c.JSON(http.StatusCreated, common.Response{
			Error:   false,
			Code:    201,
			Message: "Transaction created successfully",
			Data:    newTransaction,
		})
	}
}

func GetTransactionsByAccountId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transactionList []models.TransactionProps
		defer cancel()

		id := c.Param("id")
		cursor, err := transactions.Find(ctx, bson.M{"account_id": id})
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    404,
				Message: "Transaction not found!",
				Data:    nil,
			})
			return
		}
		if err = cursor.All(ctx, &transactionList); err != nil {
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
			Message: "Transactions found!",
			Data:    transactionList,
		})
	}
}

func GetTransactionById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction *models.TransactionProps
		defer cancel()

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := transactions.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    404,
				Message: "Transaction not found!",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Error:   false,
			Code:    200,
			Message: "Transaction found!",
			Data:    transaction,
		})
	}
}

func DeleteTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction *models.TransactionProps
		defer cancel()

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := transactions.FindOne(ctx, id).Decode(&transaction)
		if err != nil {
			c.JSON(http.StatusNotFound, common.Response{
				Error:   true,
				Code:    404,
				Message: "Account not found!",
				Data:    nil,
			})
			return
		}
		_, err = transactions.DeleteOne(ctx, id)
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
