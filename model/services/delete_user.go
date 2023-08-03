package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
)

func (ud *userDomainService) DeleteUser(userID string) *error_handle.ErrorHandle {
	logger.Info("Init deleteUser model", logger.AddJourneyTag(logger.DeleteUserJourney))

	err := ud.repository.DeleteUser(userID)
	if err != nil {
		logger.Error("Error trying to call deleteUser repository", err, logger.AddJourneyTag(logger.DeleteUserJourney))
		return error_handle.NewInternalServerError(err.Error())
	}

	logger.Info("deleteUser service executed successfully", logger.AddGenericTag("userID", userID), logger.AddJourneyTag(logger.DeleteUserJourney))
	return nil
}
