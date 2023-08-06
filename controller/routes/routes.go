package routes

import (
	"github.com/MogLuiz/go-person-api/controller"
	"github.com/MogLuiz/go-person-api/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifyTokenMiddleware, userController.CreateUser)
	r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.Login)
}
