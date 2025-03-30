package application

import (
	"api-v1/src/user/domain/entities"
	"api-v1/src/user/domain/ports"
)

type GetByIdUserUseUseCase struct {
	UserRepository ports.IUserRepository
}

func NewGetByIdUserUseCase(userRepository ports.IUserRepository) *GetByIdUserUseUseCase{
	return &GetByIdUserUseUseCase{UserRepository: userRepository}
}

func (s GetByIdUserUseUseCase) Run(id int64) (entities.User, error){
	user, err := s.UserRepository.GetByID(id)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}