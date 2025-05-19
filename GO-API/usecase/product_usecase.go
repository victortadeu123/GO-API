package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

	//get
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return  pu.repository.GetProducts()
}


	//delete
func (pu *ProductUsecase) DeleteProduct(id string) error {
	return pu.repository.DeleteProduct(id)
}

//Post
func (pu *ProductUsecase) CreateProduct(p *model.Product) error {
	return pu.repository.CreateProduct(p)
}


//PUT

func (pu *ProductUsecase) UpdateProduct(p model.Product) error {
	return pu.repository.UpdateProduct(p)
}


