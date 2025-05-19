package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}


//delete
func (r *ProductRepository) DeleteProduct(id string) error {
	query := "DELETE FROM product WHERE id = $1"
	result, err := r.connection.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

	//get
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	
	
	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	defer rows.Close() 


	return productList, nil
}

// post
func (r *ProductRepository) CreateProduct(p *model.Product) error {
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id"
	err := r.connection.QueryRow(query, p.Name, p.Price).Scan(&p.ID)
	return err
}


//PUT
// repository.go
func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	query := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3"
	_, err := pr.connection.Exec(query, product.Name, product.Price, product.ID)
	return err
}

//colocar o id apos o update, mostrando qual produto com id foi criado
