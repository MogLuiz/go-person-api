package repository

import (
	"os"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv("DATABASE_NAME", database_name)
}
