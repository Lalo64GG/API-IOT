package ports

import "api-v1/src/user/domain/entities"


type IUserRepository interface {
	Create(user entities.User) (entities.User, error)
	GetByID(id int64) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
}
