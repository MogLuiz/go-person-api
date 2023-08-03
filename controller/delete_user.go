package controller

import (
	"net/http"
	"strings"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller", logger.AddJourneyTag(logger.DeleteUserJourney))

	userID := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userID); err != nil || strings.TrimSpace(userID) == "" {
		errorMessage := error_handle.NewBadRequestError("userId is not a valid UUID")
		logger.Error("Error trying to validate userId", err, logger.AddJourneyTag(logger.FindUserByIDJourney))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := uc.service.DeleteUser(userID)
	if err != nil {
		logger.Error("Error trying to call deleteUser service", err, logger.AddJourneyTag(logger.DeleteUserJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("deleteUser controller executed successfully", logger.AddGenericTag("userID", userID), logger.AddJourneyTag(logger.DeleteUserJourney))
	c.Status(http.StatusOK)
}
