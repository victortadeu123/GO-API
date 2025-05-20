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
    query := "SELECT id, product_name, price, estoque, category, description, active FROM product order by id"
    rows, err := pr.connection.Query(query)
    if err != nil {
        fmt.Println(err)
        return nil, err // Retornar nil em vez de slice vazia para erros
    }
    defer rows.Close() // Melhor colocar o defer logo após verificar o erro

    var products []model.Product

    for rows.Next() {
        var product model.Product
        err = rows.Scan(
            &product.ID,
            &product.Name,
            &product.Price,
            &product.Estoque,
            &product.Categoria,
            &product.Descricao,
            &product.Ativo)
        
        if err != nil {
            fmt.Println(err)
            return nil, err
        }

        products = append(products, product)
    }

    // Verificar erros que podem ter ocorrido durante a iteração
    if err = rows.Err(); err != nil {
        fmt.Println(err)
        return nil, err
    }

    return products, nil
}


// post
func (r *ProductRepository) CreateProduct(p *model.Product) error {
	query := "INSERT INTO product (product_name, price, estoque, category, description, active) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := r.connection.QueryRow(query, p.Name, p.Price, p.Estoque, p.Categoria, p.Descricao, p.Ativo).Scan(&p.ID)
	return err
}


//PUT
func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	query := "UPDATE product SET product_name = $1, price = $2, estoque = $3, category = $4, description = $5, active = $6 WHERE id = $7"
	_, err := pr.connection.Exec(query, product.Name, product.Price, product.Estoque, product.Categoria, product.Descricao, product.Ativo, product.ID)
	return err
}

//colocar o id apos o update, mostrando qual produto com id foi criado
