package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUser model", logger.AddJourneyTag(logger.FindUserJourney))

	return nil, nil
}
