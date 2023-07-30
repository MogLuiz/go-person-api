package controller

import (
	"net/http"

	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/configuration/validation"
	"github.com/MogLuiz/go-person-api/src/controller/model/request"
	"github.com/MogLuiz/go-person-api/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", logger.AddJourneyTag(logger.CreateUserJourney))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, logger.AddJourneyTag(logger.CreateUserJourney))
		error := validation.ValidateUserError(err)

		c.JSON(error.Code, error)
		return
	}

	response := response.UserResponse{
		ID:    "1",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	logger.Info("user created successfully", logger.AddJourneyTag(logger.CreateUserJourney))
	c.JSON(http.StatusOK, response)
}
