package controller

import (
	"net/http"

	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/configuration/validation"
	"github.com/MogLuiz/go-person-api/controller/model/request"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/view"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", logger.AddJourneyTag(logger.CreateUserJourney))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, logger.AddJourneyTag(logger.CreateUserJourney))
		error := validation.ValidateUserError(err)

		c.JSON(error.Code, error)
		return
	}

	userDomain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call createUser service", err, logger.AddJourneyTag(logger.CreateUserJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("createdUser controller executed successfully", logger.AddGenericTag("userID", domainResult.GetID()), logger.AddJourneyTag(logger.CreateUserJourney))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
