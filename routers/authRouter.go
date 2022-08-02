package routers

import (
	"github.com/KingAnointing/go-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	router.POST("/api/signup", controllers.CreateASoldierProfile())

}
