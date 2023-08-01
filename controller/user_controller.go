package controller

import (
	"github.com/MogLuiz/go-person-api/model/services"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface services.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByEmail(c *gin.Context)
	FindUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}
