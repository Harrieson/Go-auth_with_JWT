package controllers

import (
	"go-jwt/database"
	helper "go-jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var useCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HarshPassword() {

}

func Verifypassword() {

}

func Signup() {

}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

}
