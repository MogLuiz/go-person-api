package services

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
)

func (ud *userDomainService) DeleteUser(string) *error_handle.ErrorHandle {
	logger.Info("Init deleteUser model", logger.AddJourneyTag(logger.DeleteUserJourney))

	return nil
}