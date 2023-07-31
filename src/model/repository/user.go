package repository

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/model"
	"go.mongodb.org/mongo-driver/mongo"
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
}
