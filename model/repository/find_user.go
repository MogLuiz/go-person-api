package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/model/repository/entity"
	"github.com/MogLuiz/go-person-api/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByEmail repository", logger.AddJourneyTag(logger.UpdateUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", email)
			logger.Error(errorMessage, err, logger.AddJourneyTag(logger.UpdateUserJourney))
			return nil, error_handle.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error trying to find user with email %s", email)
		logger.Error(errorMessage, err, logger.AddJourneyTag(logger.UpdateUserJourney))
		return nil, error_handle.NewInternalServerError(errorMessage)
	}

	logger.Info("findUserByEmail repository executed successfully",
		logger.AddGenericTag("userID", userEntity.ID.Hex()),
		logger.AddGenericTag("email", userEntity.Email),
		logger.AddJourneyTag(logger.UpdateUserJourney))

	return converter.ConvertEntityToDomain(userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByID repository", logger.AddJourneyTag(logger.UpdateUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with ID %s not found", id)
			logger.Error(errorMessage, err, logger.AddJourneyTag(logger.UpdateUserJourney))
			return nil, error_handle.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error trying to find user with ID %s", id)
		logger.Error(errorMessage, err, logger.AddJourneyTag(logger.UpdateUserJourney))
		return nil, error_handle.NewInternalServerError(errorMessage)
	}

	logger.Info("findUserByID repository executed successfully",
		logger.AddGenericTag("userID", userEntity.ID.Hex()),
		logger.AddJourneyTag(logger.UpdateUserJourney))

	return converter.ConvertEntityToDomain(userEntity), nil
}
