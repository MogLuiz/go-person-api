package repository

import (
	"context"
	"os"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) DeleteUser(userID string) *error_handle.ErrorHandle {
	logger.Info("Init deleteUser repository", logger.AddJourneyTag(logger.DeleteUserJourney))

	collection := ur.databaseConnection.Collection(os.Getenv(MOGODB_USER_COLLECTION))

	userIDHex, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.D{{Key: "_id", Value: userIDHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error on delete user in database", err, logger.AddJourneyTag(logger.DeleteUserJourney))
		return error_handle.NewInternalServerError(err.Error())
	}

	logger.Info("deleteUser repository executed successfully", logger.AddGenericTag("userID", userID), logger.AddJourneyTag(logger.DeleteUserJourney))
	return nil
}
