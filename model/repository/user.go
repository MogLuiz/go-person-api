package repository

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MOGODB_USER_COLLECTION = "MOGODB_USER_COLLECTION"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *error_handle.ErrorHandle)
	FindUserByEmail(email string) (model.UserDomainInterface, *error_handle.ErrorHandle)
	FindUserByID(id string) (model.UserDomainInterface, *error_handle.ErrorHandle)
	UpdateUser(userID string, userDomain model.UserDomainInterface) *error_handle.ErrorHandle
}
