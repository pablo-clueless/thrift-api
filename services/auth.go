package services

import (
	"context"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"thrift.com/m/common"
	"thrift.com/m/database"
	"thrift.com/m/models"
)

var users *mongo.Collection = database.GetCollection(database.Database, "users")

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.UserProps
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid request payload!",
				Data:    nil,
			})
			return
		}

		if !govalidator.IsEmail(user.Email) {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid email address!",
				Data:    nil,
			})
			return
		}

		// check if user already exists
		count, err := users.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Something went wrong!",
				Data:    nil,
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusConflict, common.Response{
				Error:   true,
				Code:    409,
				Message: "This email has been registered!",
				Data:    nil,
			})
			return
		}

		hash := common.HashPassword(user.Password)
		newUser := models.UserProps{
			Id:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			Email:     user.Email,
			ImageUrl:  user.ImageUrl,
			Name:      user.Name,
			Password:  hash,
			Username:  user.Username,
		}
		_, err = users.InsertOne(ctx, newUser)
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
			Message: "User created successfully!",
			Data:    nil,
		})
	}
}

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var credentials *models.AuthProps
		var user *models.UserProps
		defer cancel()

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid request payload!",
				Data:    nil,
			})
			return
		}

		err := users.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Error:   true,
				Code:    400,
				Message: "Invalid email address!",
				Data:    nil,
			})
			return
		}

		id := user.Id.Hex()
		token, err := common.GenerateToken(id)
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
			Message: "User logged in successfully!",
			Data:    map[string]interface{}{"user": user, "token": token},
		})
	}
}
