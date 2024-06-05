package product

import "github.com/evrintobing17/ecommence-REST/app/models"

type ProductUsecase interface {
	CreateProduct(product models.Product) (*models.Product, error)
	GetBySellerID(id int) (*[]models.Product, error)
	GetAll() (*[]models.Product, error)
}
