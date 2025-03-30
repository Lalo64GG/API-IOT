package ports

import "api-v1/src/product/domain/entities"

type IProductRepository interface{
	Create(product entities.Product) (entities.Product, error)
	GetAll()([]entities.Product, error)
}
