package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *error_handle.ErrorHandle {
	logger.Info("Init updateUser model", logger.AddJourneyTag(logger.UpdateUserJourney))

	err := ud.repository.UpdateUser(userID, userDomain)
	if err != nil {
		logger.Error("Error trying to call updateUser repository", err, logger.AddJourneyTag(logger.UpdateUserJourney))
		return error_handle.NewInternalServerError(err.Error())
	}

	logger.Info("updateUser service executed successfully", logger.AddGenericTag("userID", userID), logger.AddJourneyTag(logger.UpdateUserJourney))
	return nil
}
