package http

import (
	"api-v1/src/horario/application"
	"api-v1/src/horario/domain/ports"
	"api-v1/src/horario/infraestructure/adapters"
	"api-v1/src/horario/infraestructure/http/controllers"
	"log"
)

var horarioRepository ports.IHorarioRepository

func init(){
	var err error
	horarioRepository, err = adapters.NewHorarioRepositoryMysql()

	if err != nil {
		log.Fatalf("Error creating horario repository %v", err)
	}
}

func SetUpCreateController() *controllers.CreateHorarioController{
	createUseCase := application.NewCreateHorarioUseCase(horarioRepository)
	return controllers.NewCreateHorarioController(createUseCase)
}

func SetUpGetAllController() *controllers.GetAllHorarioController{
	getAllUseCase := application.NewGetAllHorarioUseCase(horarioRepository)
	return controllers.NewGetAllHorarioController(getAllUseCase)
}

func SetUpDeleteController() *controllers.DeleteHorarioController{
	deleteUseCase := application.NewDeleteHorarioUseCase(horarioRepository)
	return controllers.NewDeleteHorarioController(deleteUseCase)
}