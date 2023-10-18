package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByID service", logger.AddJourneyTag(logger.FindUserByIDJourney))

	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByEmail service", logger.AddJourneyTag(logger.FindUserByEmailJourney))

	return ud.repository.FindUserByEmail(email)

}

func (ud *userDomainService) findUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByEmail service", logger.AddJourneyTag(logger.LoginUserJourney))

	return ud.repository.FindUserByEmailAndPassword(email, password)
}
