package main

import (
	"fmt"
	"log"

	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/controller"
	"github.com/MogLuiz/go-person-api/src/controller/routes"
	"github.com/MogLuiz/go-person-api/src/model/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting API...")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	service := services.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
