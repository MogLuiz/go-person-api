package services

import (
	"testing"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/MogLuiz/go-person-api/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_FindUserByIDService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repository)

	t.Run("it should return success when exists user", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "123", "john doe", 21)

		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		returnedUser, err := service.FindUserByID(id)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, returnedUser)
	})

	t.Run("it should return error when not exists user", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserByID(id).Return(nil, error_handle.NewNotFoundError("user not found"))

		returnedUser, err := service.FindUserByID(id)

		assert.Nil(t, returnedUser)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "user not found")
	})
}
