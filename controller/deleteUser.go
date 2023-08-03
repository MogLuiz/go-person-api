package controller

import (
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

}
