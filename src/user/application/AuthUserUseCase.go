package application

import (
	"api-v1/src/user/domain/entities"
	"api-v1/src/user/domain/ports"
)

type AuthUserUseUseCase struct {
	UserRepository ports.IUserRepository
}

func NewAuthUserUseCase(userRepository ports.IUserRepository) *AuthUserUseUseCase{
	return &AuthUserUseUseCase{UserRepository: userRepository}
}

func (s AuthUserUseUseCase) Run(email string)(entities.User, error){
	user, err := s.UserRepository.GetByEmail(email)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}