package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByID service", logger.AddJourneyTag(logger.FindUserByIDJourney))

	userDomainRepository, err := ud.repository.FindUserByID(id)
	if err != nil {
		logger.Error("Error trying to call findUserByID repository", err, logger.AddJourneyTag(logger.FindUserByIDJourney))
		return nil, error_handle.NewInternalServerError(err.Error())
	}

	return userDomainRepository, nil
}

func (ud *userDomainService) FindUserByEmail(id string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByEmail service", logger.AddJourneyTag(logger.FindUserByEmailJourney))

	return nil, nil
}
