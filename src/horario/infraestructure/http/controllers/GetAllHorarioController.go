package controllers

import (
	"api-v1/src/horario/application"
	"api-v1/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllHorarioController struct {
	HorarioService *application.GetAllHorarioUseCase
}

func NewGetAllHorarioController(horarioService *application.GetAllHorarioUseCase) *GetAllHorarioController {
	return &GetAllHorarioController{
		HorarioService: horarioService,
	}
}

func (ctr *GetAllHorarioController) Run(ctx *gin.Context) {
	horarios, err := ctr.HorarioService.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success:  false,
			Message: "Error fetching horarios",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success:  true,
		Message: "Horarios fetched successfully",
		Data:    horarios,
		Error:  nil,
	})
}