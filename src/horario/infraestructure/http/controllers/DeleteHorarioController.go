package controllers

import (
	"api-v1/src/horario/application"
	"api-v1/src/shared/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteHorarioController struct {
	HorarioService *application.DeleteHorarioUseCase
}

func NewDeleteHorarioController(horarioService *application.DeleteHorarioUseCase) *DeleteHorarioController {
	return &DeleteHorarioController{
		HorarioService: horarioService,
	}
}

func (ctr *DeleteHorarioController)Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Invalid ID",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	horario, err := ctr.HorarioService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success:  false,
			Message: "Error deleting horario",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success:  true,
		Message: "Horario deleted successfully",
		Data:    horario,
		Error:  nil,
	})

}