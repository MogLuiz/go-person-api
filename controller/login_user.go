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

func (uc *userControllerInterface) Login(c *gin.Context) {
	logger.Info("Init loginUser controller", logger.AddJourneyTag(logger.LoginUserJourney))
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, logger.AddJourneyTag(logger.LoginUserJourney))
		error := validation.ValidateUserError(err)

		c.JSON(error.Code, error)
		return
	}

	userDomain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)

	domainResult, token, err := uc.service.LoginUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call loginUser service", err, logger.AddJourneyTag(logger.LoginUserJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("loginUser controller executed successfully", logger.AddGenericTag("userID", domainResult.GetID()), logger.AddJourneyTag(logger.LoginUserJourney))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
