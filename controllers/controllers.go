package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/KingAnointing/go-project/responses"
	"github.com/gin-gonic/gin"
)

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
	}
}
