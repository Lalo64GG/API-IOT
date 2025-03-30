package routes

import (
	"api-v1/src/product/infraestructure/http"
	"api-v1/src/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup){
	middlewareAuth := middlewares.JWTAuthMiddleware()
	createController := http.SetUpCreateController()
	getAllController := http.SetUpGetAllController()

	router.Use(middlewareAuth)
	router.POST("/", createController.Run)
	router.GET("/", getAllController.Run)
}