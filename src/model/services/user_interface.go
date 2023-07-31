package services

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/model"
	"github.com/MogLuiz/go-person-api/src/model/repository"
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
	FindUser(string) (*model.UserDomainInterface, *error_handle.ErrorHandle)
	DeleteUser(string) *error_handle.ErrorHandle
}
