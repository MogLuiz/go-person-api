package model

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
)

func (ud *userDomain) UpdateUser(string) *error_handle.ErrorHandle {
	logger.Info("Init updateUser model", logger.AddJourneyTag(logger.UpdateUserJourney))

	return nil
}
