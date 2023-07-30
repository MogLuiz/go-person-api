package services

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/model"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUser model", logger.AddJourneyTag(logger.FindUserJourney))

	return nil, nil
}
