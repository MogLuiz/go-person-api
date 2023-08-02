package controller

import (
	"net/http"
	"net/mail"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller", logger.AddJourneyTag(logger.FindUserByIDJourney))

	userID := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		errorMessage := error_handle.NewBadRequestError("userId is not a valid UUID")
		logger.Error("Error trying to validate userId", err, logger.AddJourneyTag(logger.FindUserByIDJourney))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(userID)
	if err != nil {
		logger.Error("Error trying to call findUserByID service", err, logger.AddJourneyTag(logger.FindUserByIDJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully", logger.AddGenericTag("userID", userDomain.GetID()), logger.AddJourneyTag(logger.FindUserByIDJourney))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", logger.AddJourneyTag(logger.FindUserByEmailJourney))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := error_handle.NewBadRequestError("userEmail is not a valid email")
		logger.Error("Error trying to validate userEmail", err, logger.AddJourneyTag(logger.FindUserByEmailJourney))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userEmail)
	if err != nil {
		if err.Code == 404 {
			logger.Error("Error 404 when FindUserByEmail service is called", err, logger.AddJourneyTag(logger.FindUserByEmailJourney))
			c.JSON(err.Code, err)
			return
		}

		logger.Error("Error trying to call FindUserByEmail service", err, logger.AddJourneyTag(logger.FindUserByEmailJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", logger.AddGenericTag("userEmail", userDomain.GetEmail()), logger.AddJourneyTag(logger.FindUserByEmailJourney))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
