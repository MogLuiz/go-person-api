package services

import (
	"testing"

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
}
