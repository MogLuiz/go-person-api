package controller

import (
	"net/http"

	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/configuration/validation"
	"github.com/MogLuiz/go-person-api/src/controller/model/request"
	"github.com/MogLuiz/go-person-api/src/model"
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

	userDomain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	err := userDomain.CreateUser()
	if err != nil {
		logger.Error("Error trying to create user", err, logger.AddJourneyTag(logger.CreateUserJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created successfully", logger.AddJourneyTag(logger.CreateUserJourney))
	c.String(http.StatusOK, "user created successfully")
}
