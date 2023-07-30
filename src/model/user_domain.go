package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/MogLuiz/go-person-api/src/configuration/error_handle"
)

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *error_handle.ErrorHandle
	UpdateUser(string) *error_handle.ErrorHandle
	FindUser(string) (*userDomain, *error_handle.ErrorHandle)
	DeleteUser(string) *error_handle.ErrorHandle
}
