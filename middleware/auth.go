package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"thrift.com/m/common"
	"thrift.com/m/database"
	"thrift.com/m/models"
)

var users *mongo.Collection = database.GetCollection(database.Database, "users")

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.UserProps

		token := c.GetHeader("Authorization")
		token = strings.Split(token, " ")[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, common.Response{
				Error:   true,
				Code:    401,
				Message: "No auth token found!",
				Data:    nil,
			})
			c.Abort()
			return
		}

		claims, err := common.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.Response{
				Error:   true,
				Code:    401,
				Message: "Invalid auth token!",
				Data:    nil,
			})
			c.Abort()
			return
		}
		id, _ := primitive.ObjectIDFromHex(claims.Id)
		err = users.FindOne(c, primitive.M{"_id": id}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.Response{
				Error:   true,
				Code:    401,
				Message: "Unauthorized!",
				Data:    nil,
			})
			c.Abort()
			return
		}

		today := time.Now().UTC().Truncate(24 * time.Hour)
		if user.LastLogin.IsZero() || user.LastLogin.Before(today) {
			user.Streak = 0
		} else {
			user.Streak++
		}

		user.LastLogin = today
		_, errs := users.UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": user})
		if errs != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Error:   true,
				Code:    500,
				Message: "Internal Server Error!",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
