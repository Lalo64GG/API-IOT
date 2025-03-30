package ports

import "api-v1/src/horario/domain/entities"

type IHorarioRepository interface {
	Create(horario entities.Horario) (entities.Horario, error)
	GetAll() ([]entities.Horario, error)
	Delete(id int64) (bool, error)
}