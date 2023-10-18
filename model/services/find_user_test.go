package services

import (
	"testing"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/test/mocks"
	"go.uber.org/mock/gomock"
)

func TestFindUserByIDService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	id := "123"

	repository := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repository)

	repository.EXPECT().FindUserByID(id).Return(model.NewUserDomain("test@test.com", "123", "john doe", 21), nil)

	user, err := service.FindUserByID(id)

	if err != nil {
		t.Errorf("Error should be nil, but got %v", err)
	}
	if user == nil {
		t.Errorf("User should not be nil")
	}
	if user.GetEmail() != "test@test.com" {
		t.Errorf("User email should be test@test.com but got %v", user.GetEmail())
	}

	repository2 := mocks.NewMockUserRepository(control)
	service2 := NewUserDomainService(repository2)

	repository2.EXPECT().FindUserByID(id).Return(nil, error_handle.NewNotFoundError("User not found"))

	user2, err2 := service2.FindUserByID(id)

	if err2 == nil {
		t.Errorf("Error should not be nil")
	}
	if user2 != nil {
		t.Errorf("User should be nil")
	}
	if err2.Message != "User not found" {
		t.Errorf("Error message should be User not found but got %v", err2.Message)
	}
}
