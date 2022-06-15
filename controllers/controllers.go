package controllers

import (
	"net/http"

	"github.com/KingAnointing/go-project/responses"
	"github.com/gin-gonic/gin"
)

func Greeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "sucess", Data: map[string]interface{}{"data": "Welcome to my personal skill test on gin, monogoDB & golang CRUD API"}})
	}
}
