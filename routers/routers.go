package routers

import (
	"github.com/KingAnointing/go-project/controllers"
	"github.com/gin-gonic/gin"
)

func Router()  {
	router := gin.Default()

	router.GET("/api/",controllers.Greeter())
	router.POST("/api/soldier",controllers.CreateSoldier())

	router.Run(":8080")
}