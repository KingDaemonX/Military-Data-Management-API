package routers

import (
	"github.com/KingAnointing/go-project/controllers"
	"github.com/gin-gonic/gin"
)

func Router( router *gin.Engine) {
	router.GET("/api/", controllers.Greeter())
	router.POST("/api/soldiers", controllers.CreateASoldierProfile())
	router.GET("/api/soldiers/:userId", controllers.GetOneSoldierProfile())
	router.GET("/api/soldiers", controllers.GetAllSoldierProfile())
	router.PUT("/api/soldiers/:userId", controllers.UpdateSoldierProfile())
	router.DELETE("/api/soldiers/:userId", controllers.DeleteASoldierProfile())
	router.DELETE("/api/deleteall", controllers.DeleteAllSoldierProfile())
}
