package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init createUser model", logger.AddJourneyTag(logger.CreateUserJourney))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call createUser repository", err, logger.AddJourneyTag(logger.CreateUserJourney))
		return nil, error_handle.NewInternalServerError(err.Error())
	}

	logger.Info("createdUser service executed successfully", logger.AddGenericTag("userID", userDomainRepository.GetID()), logger.AddJourneyTag(logger.CreateUserJourney))
	return userDomainRepository, nil
}
