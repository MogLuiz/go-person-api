package view

import (
	"github.com/MogLuiz/go-person-api/src/controller/model/response"
	"github.com/MogLuiz/go-person-api/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
