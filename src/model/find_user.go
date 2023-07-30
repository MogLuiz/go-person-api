package model

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
)

func (ud *userDomain) FindUser(string) (*userDomain, *error_handle.ErrorHandle) {
	logger.Info("Init findUser model", logger.AddJourneyTag(logger.FindUserJourney))

	return nil, nil
}
