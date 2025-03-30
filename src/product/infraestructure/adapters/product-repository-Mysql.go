package adapaters

import (
	"api-v1/src/database"
	"api-v1/src/product/domain/entities"
	"database/sql"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql() (*ProductRepositoryMysql, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	return &ProductRepositoryMysql{DB: db}, nil
}


func (r *ProductRepositoryMysql) Create(product entities.Product) (entities.Product, error) {
	query := "INSERT INTO products (Name, Fecha_Adquisicion) VALUES (?,?)"
    result, err := r.DB.Exec(query, product.Name, product.Fecha_Adquisicion)

    if err != nil {
        return entities.Product{}, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return entities.Product{}, err
    }

    product.ID = int(id)

    return product, nil
}

func (r *ProductRepositoryMysql) GetAll() ([]entities.Product, error) {
	query := "SELECT ID, Name, Fecha_Adquisicion FROM products"
	stmt, err := r.DB.Prepare(query)
	
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Fecha_Adquisicion)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil

}