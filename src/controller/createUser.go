package controller

import (
	"fmt"

	"github.com/MogLuiz/go-person-api/src/configuration/error_logger"
	"github.com/MogLuiz/go-person-api/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		error := error_logger.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		c.JSON(error.Code, error)
		return
	}

	fmt.Println(userRequest)
}
