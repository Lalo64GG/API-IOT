package application

import (
	"api-v1/src/user/application/services"
	"api-v1/src/user/domain/entities"
	"api-v1/src/user/domain/ports"
)

type CreateUserUseCase struct {
	UserRepository ports.IUserRepository
	EncryptService services.EncryptService
}

func NewCreateUserUseCase(userRepository ports.IUserRepository, encryptService services.EncryptService ) *CreateUserUseCase{
	return &CreateUserUseCase{UserRepository: userRepository, EncryptService: encryptService}
}

func (s CreateUserUseCase) Run(Name, Email, Password string) (entities.User, error){
	encryptedPass, err := s.EncryptService.Encrypt([]byte(Password))

	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		Name:     Name,
		Email:    Email,
		Password: encryptedPass,
	}

	newUser, err := s.UserRepository.Create(user)
	
	if err != nil {
		return entities.User{}, err
	}

	return newUser, nil
}