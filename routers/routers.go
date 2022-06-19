package routers

import (
	"github.com/KingAnointing/go-project/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/api/", controllers.Greeter())
	router.POST("/api/soldier", controllers.CreateASoldierProfile())
	router.GET("/api/soldier/:userId", controllers.GetOneSoldierProfile())

	router.Run(":8080")
}
