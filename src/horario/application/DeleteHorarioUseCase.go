package application

import "api-v1/src/horario/domain/ports"

type DeleteHorarioUseCase struct{
	HorarioRepository ports.IHorarioRepository
}

func NewDeleteHorarioUseCase(horarioRepository ports.IHorarioRepository) *DeleteHorarioUseCase{
	return &DeleteHorarioUseCase{HorarioRepository: horarioRepository}
}

func (h *DeleteHorarioUseCase) Run(id int64) (bool, error){
	deleted, err := h.HorarioRepository.Delete(id)
	
	if err != nil {
		return false, err
	}

	return deleted, nil
}