package repository

import (
	"context"
	"os"

	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/configuration/logger"
	"github.com/MogLuiz/go-person-api/src/model"
)

const (
	MOGODB_USER_COLLECTION = "MOGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init createUser repository", logger.AddJourneyTag(logger.CreateUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	value, err := userDomain.GetJSONValue()
	if err != nil {
		logger.Error("Error on get json value", err, logger.AddJourneyTag(logger.CreateUserJourney))
		return nil, error_handle.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error on insert user in database", err, logger.AddJourneyTag(logger.CreateUserJourney))
		return nil, error_handle.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
