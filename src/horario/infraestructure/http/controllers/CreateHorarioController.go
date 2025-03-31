package controllers

import (
	"api-v1/src/horario/application"
	"api-v1/src/horario/infraestructure/http/request"
	"api-v1/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateHorarioController struct{
	HorarioService *application.CreateHorarioUseCase
	Validate validator.Validate
}

func NewCreateHorarioController(horarioService *application.CreateHorarioUseCase) *CreateHorarioController{
	return &CreateHorarioController{
		HorarioService: horarioService,
		Validate: *validator.New(),
	}
}


func (ctr *CreateHorarioController) Run(ctx *gin.Context){
	var req request.CreateHorarioRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Invalid request",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	if err := ctr.Validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Invalid request",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	horario, err := ctr.HorarioService.Run(req.Hour, req.Minute)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success:  false,
			Message: "Error creating horario",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success:  true,
		Message: "Horario created successfully",
		Data:    horario,
		Error:  nil,
	})
}