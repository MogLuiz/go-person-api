package services

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init createUser model", logger.AddJourneyTag(logger.CreateUserJourney))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error on create user in repository", err, logger.AddJourneyTag(logger.CreateUserJourney))
		return nil, err
	}

	return userDomainRepository, nil
}
