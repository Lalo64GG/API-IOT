package controllers

import (
	"api-v1/src/shared/responses"
	"api-v1/src/user/application"
	"api-v1/src/user/infraestructure/http/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserController struct{
	UserService *application.CreateUserUseCase
	Validate *validator.Validate
}

func NewCreateUserController(userService *application.CreateUserUseCase) *CreateUserController{
	return &CreateUserController{
		UserService: userService,
		Validate: validator.New(),
	}
}

func (ctr *CreateUserController) Run(ctx *gin.Context){
	var req request.CreateUserRequest

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

	user, err := ctr.UserService.Run(req.Name, req.Email, req.Password)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success:  false,
			Message: "Error creating user",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success:  true,
		Message: "User created successfully",
		Data:    user,
		Error:  nil,
	})
}