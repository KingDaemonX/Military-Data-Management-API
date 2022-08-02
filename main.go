package main

import (
	"fmt"

	"github.com/KingAnointing/go-project/configs"
	"github.com/KingAnointing/go-project/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to my personal pet project")
	fmt.Println("I'm building it with golang, gin-gonic & mongodb while using other dependency")

	configs.ConnectDB()
	fmt.Println("Database is Working Perfectly :)")

	// router
	router := gin.New()
	router.Use(gin.Logger())
	routers.Router(router)
	router.Run(":8080")
}
