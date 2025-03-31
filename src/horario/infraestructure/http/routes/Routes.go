package routes

import (
	"api-v1/src/horario/infraestructure/http"
	"api-v1/src/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func HorarioRoutes(router *gin.RouterGroup){
	middlewareAuth := middlewares.JWTAuthMiddleware()
	getAllController := http.SetUpGetAllController()
	getByIdController := http.SetUpDeleteController()
	createController := http.SetUpCreateController()

	router.Use(middlewareAuth)
	router.POST("/", createController.Run)
	router.GET("/", getAllController.Run)
	router.DELETE("/:id", getByIdController.Run)
	
}