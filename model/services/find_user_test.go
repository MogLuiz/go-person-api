package services

import (
	"math/rand"
	"strconv"
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
		userDomain.SetID(id)

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

func TestUserDomainService_FindUserByEmailService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repository)

	t.Run("it should return success when exists user", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"

		userDomain := model.NewUserDomain(email, "123", "john doe", 21)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		returnedUser, err := service.FindUserByEmail(email)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, returnedUser)
	})

	t.Run("it should return error when not exists user", func(t *testing.T) {
		email := "test@failure.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, error_handle.NewNotFoundError("user not found"))

		returnedUser, err := service.FindUserByEmail(email)

		assert.Nil(t, returnedUser)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailAndPasswordService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repository := mocks.NewMockUserRepository(control)
	service := &userDomainService{repository}

	t.Run("it should return success when exists user", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@success.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(email, password, "john doe", 21)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		returnedUser, err := service.findUserByEmailAndPassword(email, password)

		assert.Nil(t, err)
		assert.Equal(t, userDomain, returnedUser)
	})

	t.Run("it should return error when not exists user", func(t *testing.T) {
		email := "test@failure.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, error_handle.NewNotFoundError("user not found"))

		returnedUser, err := service.findUserByEmailAndPassword(email, password)

		assert.Nil(t, returnedUser)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "user not found")
	})
}
