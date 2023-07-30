package services

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *error_handle.ErrorHandle
	UpdateUser(string, model.UserDomainInterface) *error_handle.ErrorHandle
	FindUser(string) (*model.UserDomainInterface, *error_handle.ErrorHandle)
	DeleteUser(string) *error_handle.ErrorHandle
}
