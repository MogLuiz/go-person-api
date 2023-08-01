package main

import (
	"github.com/MogLuiz/go-person-api/src/controller"
	"github.com/MogLuiz/go-person-api/src/model/repository"
	"github.com/MogLuiz/go-person-api/src/model/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(databaseConnection *mongo.Database) controller.UserControllerInterface {
	repository := repository.NewUserRepository(databaseConnection)
	service := services.NewUserDomainService(repository)
	userController := controller.NewUserControllerInterface(service)

	return userController
}
