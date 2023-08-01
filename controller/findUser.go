package controller

import (
	"net/http"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/view"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller", logger.AddJourneyTag(logger.FindUserByIDJourney))

	userID := c.Param("userId")

	if _, err := uuid.Parse(userID); err != nil {
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

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {}
