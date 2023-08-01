package services

import (
	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		repository: userRepository,
	}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *error_handle.ErrorHandle)
	UpdateUser(string, model.UserDomainInterface) *error_handle.ErrorHandle
	FindUserByID(id string) (model.UserDomainInterface, *error_handle.ErrorHandle)
	FindUserByEmail(email string) (model.UserDomainInterface, *error_handle.ErrorHandle)
	DeleteUser(string) *error_handle.ErrorHandle
}
