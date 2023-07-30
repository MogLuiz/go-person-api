package controller

import (
	"github.com/MogLuiz/go-person-api/src/configuration/error_logger"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := error_logger.NewBadRequestError("Teste de erro")
	c.JSON(err.Code, err)
}
