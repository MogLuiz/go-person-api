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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByEmail repository", logger.AddJourneyTag(logger.FindUserByEmailJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with email %s not found", email)
			logger.Error(errorMessage, err, logger.AddJourneyTag(logger.FindUserByEmailJourney))
			return nil, error_handle.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error trying to find user with email %s", email)
		logger.Error(errorMessage, err, logger.AddJourneyTag(logger.FindUserByEmailJourney))
		return nil, error_handle.NewInternalServerError(errorMessage)
	}

	logger.Info("findUserByEmail repository executed successfully",
		logger.AddGenericTag("userID", userEntity.ID.Hex()),
		logger.AddGenericTag("email", userEntity.Email),
		logger.AddJourneyTag(logger.FindUserByEmailJourney))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByID repository", logger.AddJourneyTag(logger.FindUserByIDJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectID}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User with ID %s not found", id)
			logger.Error(errorMessage, err, logger.AddJourneyTag(logger.FindUserByIDJourney))
			return nil, error_handle.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error trying to find user with ID %s", id)
		logger.Error(errorMessage, err, logger.AddJourneyTag(logger.FindUserByIDJourney))
		return nil, error_handle.NewInternalServerError(errorMessage)
	}

	logger.Info("findUserByID repository executed successfully",
		logger.AddGenericTag("userID", userEntity.ID.Hex()),
		logger.AddJourneyTag(logger.FindUserByIDJourney))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *error_handle.ErrorHandle) {
	logger.Info("Init findUserByEmailAndPassword repository", logger.AddJourneyTag(logger.LoginUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage, err, logger.AddJourneyTag(logger.LoginUserJourney))
			return nil, error_handle.NewForbiddenError(errorMessage)
		}

		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, err, logger.AddJourneyTag(logger.LoginUserJourney))
		return nil, error_handle.NewInternalServerError(errorMessage)
	}

	logger.Info("findUserByEmailAndPassword repository executed successfully",
		logger.AddGenericTag("userID", userEntity.ID.Hex()),
		logger.AddGenericTag("email", userEntity.Email),
		logger.AddJourneyTag(logger.LoginUserJourney))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
