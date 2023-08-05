package model

import (
	"fmt"
	"os"
	"time"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *error_handle.ErrorHandle) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", error_handle.NewInternalServerError(fmt.Sprintf("Error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
