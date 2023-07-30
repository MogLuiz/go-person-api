package model

import "github.com/MogLuiz/go-person-api/src/configuration/error_handle"

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *error_handle.ErrorHandle
	UpdateUser(string, UserDomain) *error_handle.ErrorHandle
	FindUser(string) (*UserDomain, *error_handle.ErrorHandle)
	DeleteUser(string) *error_handle.ErrorHandle
}
