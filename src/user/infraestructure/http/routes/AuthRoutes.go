package routes

import (
	"api-v1/src/user/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup){
	createController := http.SetUpCreateController()
	authController := http.AuthController()

	router.POST("/", createController.Run)
	router.POST("/auth", authController.Run)
}