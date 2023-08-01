package repository

import (
	"context"
	"os"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MOGODB_USER_COLLECTION = "MOGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init createUser repository", logger.AddJourneyTag(logger.CreateUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error on insert user in database", err, logger.AddJourneyTag(logger.CreateUserJourney))
		return nil, error_handle.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("createdUser repository executed successfully", logger.AddGenericTag("userID", value.ID.Hex()), logger.AddJourneyTag(logger.CreateUserJourney))
	return converter.ConvertEntityToDomain(value), nil
}
