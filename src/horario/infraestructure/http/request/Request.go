package request

type CreateHorarioRequest struct {
	Minute string `json:"minute" validate:"required"`
	Hour   string `json:"hour" validate:"required"`
}