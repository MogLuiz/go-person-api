package model

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MogLuiz/go-person-api/configuration/error_handle"
	"github.com/MogLuiz/go-person-api/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	secret         = os.Getenv(JWT_SECRET_KEY)
)

func (ud *userDomain) GenerateToken() (string, *error_handle.ErrorHandle) {
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

func VerifyToken(jwtToken string) (UserDomainInterface, *error_handle.ErrorHandle) {
	token, err := jwt.Parse(RemoveBearerPrefix(jwtToken), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, error_handle.NewBadRequestError("invalid token")
	})
	if err != nil {
		return nil, error_handle.NewUnauthorizedError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, error_handle.NewUnauthorizedError("invalid token")
	}

	return &userDomain{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil

}

func VerifyTokenMiddleware(c *gin.Context) {
	token := RemoveBearerPrefix(c.GetHeader("Authorization"))
	if token == "" {
		c.JSON(http.StatusUnauthorized, error_handle.NewUnauthorizedError("missing token"))
		c.Abort()
		return
	}

	user, err := VerifyToken(token)
	if err != nil {
		c.JSON(err.Code, err)
		c.Abort()
		return
	}

	logger.Info(fmt.Sprintf("User authenticated %#v", user))

	c.Next()
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
