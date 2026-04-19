package repository

import (
	"database/sql"
	"fmt"
	"korp/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, code, description, balance FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var product model.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.CODE,
			&product.DESCRIPTION,
			&product.BALANCE)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, product)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO products" +
		"(code, description, balance)" +
		" VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.CODE, product.DESCRIPTION, product.BALANCE).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) EditProduct(id int, product model.Product) (model.Product, error) {

	query, err := pr.connection.Prepare("UPDATE products SET code = $1, description = $2, balance = $3 WHERE id = $4 RETURNING id")
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	err = query.QueryRow(product.CODE, product.DESCRIPTION, product.BALANCE, id).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}

		return model.Product{}, err
	}

	newProduct := model.Product{
		ID:          id,
		CODE:        product.CODE,
		DESCRIPTION: product.DESCRIPTION,
		BALANCE:     product.BALANCE}

	query.Close()

	return newProduct, nil

}
