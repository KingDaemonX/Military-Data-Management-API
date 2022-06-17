package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/KingAnointing/go-project/configs"
	"github.com/KingAnointing/go-project/models"
	"github.com/KingAnointing/go-project/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var collections *mongo.Collection = configs.DbCollection()
var validate = validator.New()

// greeter function
func Greeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "sucess", Data: map[string]interface{}{"data": "Welcome to my personal skill test on gin, monogoDB & golang CRUD API"}})
	}
}

// create function
func CreateSoldier() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cbg, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var soldier models.Army

		// validate json response
		if err := ctx.BindJSON(&soldier); err != nil {
			ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// validate for required field and validation rule giving
		if validateErr := validate.Struct(&soldier); validateErr != nil {
			ctx.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validateErr.Error()}})
			return
		}

		result, err := collections.InsertOne(cbg, soldier)
	}
}
