package http

import (
	"api-v1/src/product/application"
	"api-v1/src/product/domain/ports"
	"api-v1/src/product/infraestructure/adapters"
	"api-v1/src/product/infraestructure/http/controller"
	"log"
)

var productRepository ports.IProductRepository

func init(){
	var err error
	productRepository, err = adapaters.NewProductRepositoryMysql()
	if err != nil {
		log.Fatalf("Error creating product repository %v", err)
	}

}

func SetUpCreateController() *controller.CreateProductController{
	createUseCase := application.NewCreateProductUseCase(productRepository)
	return controller.NewCreateProductController(createUseCase)
}

func SetUpGetAllController() *controller.GetAllProductController{
	getAllUseCase := application.NewGetAllProductUseCase(productRepository)
	return controller.NewGetAllProductController(getAllUseCase)
}