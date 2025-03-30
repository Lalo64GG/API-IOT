package application

import (
	"api-v1/src/horario/domain/entities"
	"api-v1/src/horario/domain/ports"
)

type GetAllHorarioUseCase struct {
	HorarioRepository ports.IHorarioRepository
}

func NewGetAllHorarioUseCase(horarioRepository ports.IHorarioRepository) *GetAllHorarioUseCase{
	return &GetAllHorarioUseCase{HorarioRepository: horarioRepository}
}

func (h *GetAllHorarioUseCase) Run()([]entities.Horario, error){
	horarios, err := h.HorarioRepository.GetAll()

	if err != nil {
		return []entities.Horario{}, err
	}

	return horarios, nil
}