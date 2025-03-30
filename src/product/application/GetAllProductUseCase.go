package application

import (
	"api-v1/src/product/domain/entities"
	"api-v1/src/product/domain/ports"
)

type GetAllProductUseCase struct {
	ProductRepository ports.IProductRepository
}

func NewGetAllProductUseCase(productRepository ports.IProductRepository) *GetAllProductUseCase {
	return &GetAllProductUseCase{ProductRepository: productRepository}
}

func (p *GetAllProductUseCase) Run() ([]entities.Product, error) {
	users, err := p.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}