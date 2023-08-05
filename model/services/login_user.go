package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
)

func (ud *userDomainService) LoginUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *error_handle.ErrorHandle) {
	logger.Info("Init loginUser service", logger.AddJourneyTag(logger.LoginUserJourney))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPassword(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		logger.Error("Error trying to call findUserByEmailAndPassword service", err, logger.AddJourneyTag(logger.LoginUserJourney))
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		logger.Error(err.Message, err, logger.AddJourneyTag(logger.LoginUserJourney))
		return nil, "", err
	}

	logger.Info("loginUser service executed successfully", logger.AddGenericTag("userID", user.GetID()), logger.AddJourneyTag(logger.LoginUserJourney))
	return user, token, nil
}
