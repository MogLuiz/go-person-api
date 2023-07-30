package model

import (
	"fmt"

	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
)

func (ud *UserDomain) CreateUser(UserDomain) *error_handle.ErrorHandle {
	logger.Info("Init createUser model", logger.AddJourneyTag(logger.CreateUserJourney))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
