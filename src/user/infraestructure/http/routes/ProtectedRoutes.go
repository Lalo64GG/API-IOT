package routes

import (
	"api-v1/src/shared/middlewares"
	"api-v1/src/user/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(router *gin.RouterGroup) {
	middlewareAuth := middlewares.JWTAuthMiddleware()
	getByIdController := http.GetByIdController()

	router.Use(middlewareAuth)
	router.GET("/:id", getByIdController.Run)
}