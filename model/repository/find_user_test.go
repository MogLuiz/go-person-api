package repository

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/MogLuiz/go-person-api/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("it should returns success when is sended a valid email", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBsonD(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)
		userDomain, err := repository.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)

		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDB.Run("it should throws error when mongodb returns error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)
		userDomain, err := repository.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDB.Run("it should returns user not found error", func(mt *mtest.T) {
		const testEmail = "test@teste.com"

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)
		userDomain, err := repository.FindUserByEmail(testEmail)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

		assert.EqualValues(t, err.Code, http.StatusNotFound)
		assert.EqualValues(t, err.Message, fmt.Sprintf("User with email %s not found", testEmail))
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collection_name)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("it should returns success when is sended a valid id", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBsonD(userEntity)))

		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)
		userDomain, err := repository.FindUserByID(string(userEntity.ID.Hex()))

		assert.Nil(t, err)

		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDB.Run("it should throws error when mongodb returns error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)
		userDomain, err := repository.FindUserByID("3")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDB.Run("it should returns user not found error", func(mt *mtest.T) {
		const testID = "invalid_ID"

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		databaseMock := mt.Client.Database(database_name)

		repository := NewUserRepository(databaseMock)
		userDomain, err := repository.FindUserByID(testID)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)

		assert.EqualValues(t, err.Code, http.StatusNotFound)
		assert.EqualValues(t, err.Message, fmt.Sprintf("User with ID %s not found", testID))
	})
}

func convertEntityToBsonD(entity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: entity.ID},
		{Key: "email", Value: entity.Email},
		{Key: "password", Value: entity.Password},
		{Key: "name", Value: entity.Name},
		{Key: "age", Value: entity.Age},
	}
}
