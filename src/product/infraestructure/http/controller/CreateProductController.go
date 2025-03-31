package controller

import (
	"api-v1/src/product/application"
	"api-v1/src/product/infraestructure/http/request"
	"api-v1/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateProductController struct {
	productService *application.CreateProductUseCase
	Validate *validator.Validate
}

func NewCreateProductController(productService *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{
		productService: productService,
		Validate: validator.New(),
	}
}

func (ctr *CreateProductController) Run(ctx *gin.Context){
	var req request.CreateProductRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Invalid request",
			Error:  err.Error(),
			Data: nil,
		})
		return
	}

	if err := ctr.Validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Invalid request",
			Error:  err.Error(),
			Data: nil,
		})
		return
	}

	user, err := ctr.productService.Run(req.Name, req.Fecha_Adquisicion)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success:  false,
			Message: "Error creating product",
			Error:  err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success:  true,
		Message: "Product created successfully",
		Data:    user,
		Error:  nil,
	})


}