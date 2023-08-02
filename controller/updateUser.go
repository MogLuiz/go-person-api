package controller

import (
	"net/http"
	"strings"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/configuration/validation"
	"github.com/MogLuiz/go-person-api/controller/model/request"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller", logger.AddJourneyTag(logger.UpdateUserJourney))
	var userRequest request.UserUpdateRequest

	userID := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userID); err != nil || strings.TrimSpace(userID) == "" {
		errorMessage := error_handle.NewBadRequestError("userId is not a valid UUID")
		logger.Error("Error trying to validate userId", err, logger.AddJourneyTag(logger.FindUserByIDJourney))

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, logger.AddJourneyTag(logger.UpdateUserJourney))
		error := validation.ValidateUserError(err)

		c.JSON(error.Code, error)
		return
	}

	userDomain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)
	err := uc.service.UpdateUser(userID, userDomain)
	if err != nil {
		logger.Error("Error trying to call updateUser service", err, logger.AddJourneyTag(logger.UpdateUserJourney))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("updateUser controller executed successfully", logger.AddGenericTag("userID", userID), logger.AddJourneyTag(logger.UpdateUserJourney))
	c.Status(http.StatusOK)
}
