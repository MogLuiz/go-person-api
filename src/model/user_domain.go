package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
)

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *error_handle.ErrorHandle
	UpdateUser(string) *error_handle.ErrorHandle
	FindUser(string) (*UserDomain, *error_handle.ErrorHandle)
	DeleteUser(string) *error_handle.ErrorHandle
}
