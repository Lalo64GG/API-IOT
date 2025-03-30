package controllers

import (
	"api-v1/src/shared/middlewares"
	"api-v1/src/shared/responses"
	"api-v1/src/user/application"
	"api-v1/src/user/infraestructure/http/controllers/helper"
	"api-v1/src/user/infraestructure/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{
	AuthService *application.AuthUserUseUseCase
	EncryptHelper *helper.EncryptHelper
}

func NewAuthController(authService *application.AuthUserUseUseCase) *AuthController{
	return &AuthController{AuthService: authService}
}

func (ctr *AuthController) Run(ctx *gin.Context){
	var req request.AuthRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Llene todos los campos",
			Data:    nil,
			Error:  err,
		})
		return
	}

	user, err := ctr.AuthService.Run(req.Email)

	
	if err != nil {
		switch err.Error(){
		case "sql: no rows in result set":
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "El email no existe",
                Error: err.Error(),
                Data: nil,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false,
				Message: "Error al iniciar sesión",
                Error: err.Error(),
                Data: nil,
			})
		}

		return
	}

	if err := ctr.EncryptHelper.Compare(user.Password, []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Contraseña incorrecta",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(user.ID), user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al generar el token",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Inicio de sesión exitoso",
		Data: map[string]interface{}{
			"token": token,
		},
	})

}