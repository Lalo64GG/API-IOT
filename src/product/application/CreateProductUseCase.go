package application

import (
	"api-v1/src/product/domain/entities"
	"api-v1/src/product/domain/ports"
)

type CreateProductUseCase struct {
	ProductRepository ports.IProductRepository
}

func NewCreateProductUseCase(productRepository ports.IProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (p *CreateProductUseCase) Run(Name, Fecha_Adquisicion string ) (entities.Product, error) {
    	product := entities.Product{
		Name:             Name,
		Fecha_Adquisicion: Fecha_Adquisicion,
	}

	newProduct, err := p.ProductRepository.Create(product)

	if err != nil {
		return entities.Product{}, err
	}

	return newProduct, nil
}