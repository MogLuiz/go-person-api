package services

import (
	"fmt"

	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *error_handle.ErrorHandle {
	logger.Info("Init createUser model", logger.AddJourneyTag(logger.CreateUserJourney))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
