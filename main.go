package main

import (
	"fmt"

	"github.com/KingAnointing/go-project/configs"
	"github.com/KingAnointing/go-project/routers"
)

func main() {
	fmt.Println("Welcome to my personal pet project")
	fmt.Println("I'm building it with golang, gin-gonic & mongodb while using other dependency")

	configs.ConnectDB()
	fmt.Println("Database is Working Perfectly :)")
	routers.Router()
}