package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restourant-management/controlles"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
