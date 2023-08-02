package repository

import (
	"context"
	"os"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(userID string, userDomain model.UserDomainInterface) *error_handle.ErrorHandle {
	logger.Info("Init updateUser repository", logger.AddJourneyTag(logger.UpdateUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	value := converter.ConvertDomainToEntity(userDomain)
	userIDHex, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.D{{Key: "_id", Value: userIDHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error on update user in database", err, logger.AddJourneyTag(logger.UpdateUserJourney))
		return error_handle.NewInternalServerError(err.Error())
	}

	logger.Info("updateUser repository executed successfully", logger.AddGenericTag("userID", userID), logger.AddJourneyTag(logger.UpdateUserJourney))
	return nil
}
