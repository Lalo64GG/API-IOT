package server

import (
	"api-v1/src/config"
	"api-v1/src/database"
	product "api-v1/src/product/infraestructure/http/routes"
	horario "api-v1/src/horario/infraestructure/http/routes"
	"api-v1/src/user/infraestructure/http/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	http     string
	port     string
	httpAddr string
}

func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine:   gin.New(),
		http:     http,
		port:     port,
		httpAddr: http + ":" + port,
	}

	database.Connect()
	srv.engine.Use(config.ConfigurationCors())

	srv.engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv.engine.Use(gin.Logger())
	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes(){
	userRoutes := s.engine.Group("/v1/user")
	productRoutes := s.engine.Group("/v1/product")
	horarioRoutes := s.engine.Group("/v1/horario")

	routes.UserRoutes(userRoutes)
	product.ProductRoutes(productRoutes)
	horario.HorarioRoutes(horarioRoutes)
}

func (s *Server) Run() error {
	log.Println("Starting server on " + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}