package http

import (
	"api-v1/src/user/application"
	"api-v1/src/user/application/services"
	"api-v1/src/user/domain/ports"
	"api-v1/src/user/infraestructure/adapters"
	"api-v1/src/user/infraestructure/http/controllers"
	"api-v1/src/user/infraestructure/http/controllers/helper"
	"log"
)

var (
	userRepository ports.IUserRepository
	encryptHelper services.EncryptService
)

func init(){
	var err error
	userRepository, err = adapters.NewUserRepositoryMysql()
	if err != nil {
		log.Fatalf("Error creating user repository %v", err)
	}

	encryptHelper, err = helper.NewEncryptHelper()
	if err != nil {
		log.Fatalf("Error creating encrypt service %v", err)
	}

}

func SetUpCreateController() *controllers.CreateUserController{
	createUseCase := application.NewCreateUserUseCase(userRepository, encryptHelper)
	return controllers.NewCreateUserController(createUseCase)
}

func AuthController() *controllers.AuthController{
	authUseCase := application.NewAuthUserUseCase(userRepository)
	return controllers.NewAuthController(authUseCase)
}

func GetByIdController() *controllers.GetByIdController{
	getByIdUseCase := application.NewGetByIdUserUseCase(userRepository)
    return controllers.NewGetByIdController(getByIdUseCase)
}