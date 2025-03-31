package controller

import (
	"api-v1/src/product/application"
	"api-v1/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	ProductService *application.GetAllProductUseCase
}

func NewGetAllProductController(productService *application.GetAllProductUseCase) *GetAllProductController {
    return &GetAllProductController{ProductService: productService}
}

func (ctr *GetAllProductController)Run(ctx *gin.Context){
	products, err := ctr.ProductService.Run()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success:  false,
			Message: "Error fetching products",
			Data:    nil,
			Error:  err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, responses.Response{
		Success:  true,
		Message: "Products fetched successfully",
		Data:    products,
		Error:  nil,
	})
}
