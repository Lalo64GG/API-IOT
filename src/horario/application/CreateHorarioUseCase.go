package application

import (
	"api-v1/src/horario/domain/entities"
	"api-v1/src/horario/domain/ports"
)

type CreateHorarioUseCase struct {
	HorarioRepository ports.IHorarioRepository
}

func NewCreateHorarioUseCase(horarioRepository ports.IHorarioRepository) *CreateHorarioUseCase{
	return &CreateHorarioUseCase{HorarioRepository: horarioRepository}
}

func (h *CreateHorarioUseCase) Run(Minute, Hour string) (entities.Horario, error){
	horario := entities.Horario{
		Minute: Minute,
		Hour:   Hour,
	}

	newHorario, err := h.HorarioRepository.Create(horario)
	if err != nil {
		return entities.Horario{}, err
	}

	return newHorario, nil
}