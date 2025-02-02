package routes

import (
	"github.com/HumCoding/meu-primeiro-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEMAIL)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("updateUser/:userId", controller.UpdateUser)
	r.DELETE("deleteUser/:userId", controller.DeleteUser)
}
