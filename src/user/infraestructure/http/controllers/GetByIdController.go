package controllers

import (
	"api-v1/src/shared/responses"
	"api-v1/src/user/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByIdController struct{
	UserService *application.GetByIdUserUseUseCase
}

func NewGetByIdController(userService *application.GetByIdUserUseUseCase) *GetByIdController{
    return &GetByIdController{UserService: userService}
}

func (ctr *GetByIdController)Run(ctx *gin.Context){
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Invalid ID",
			Data:    nil,
			Error:  err,
		})
        return
	}

	user, err := ctr.UserService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success:  false,
			Message: "Error fetching user",
			Data:    nil,
			Error:  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success:  true,
		Message: "User fetched successfully",
		Data:    user,
		Error:  nil,
	})
}