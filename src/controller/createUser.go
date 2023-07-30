package controller

import (
	"fmt"

	"github.com/MogLuiz/go-person-api/src/configuration/validation"
	"github.com/MogLuiz/go-person-api/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		error := validation.ValidateUserError(err)

		c.JSON(error.Code, error)
		return
	}

	fmt.Println(userRequest)
}
