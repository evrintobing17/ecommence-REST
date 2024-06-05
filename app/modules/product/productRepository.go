package product

import "github.com/evrintobing17/ecommence-REST/app/models"

type ProductRepository interface {
	GetBySellerID(id int) (*[]models.Product, error)
	GetByID(id int) (*models.Product, error)
	GetAll() (*[]models.Product, error)
	Insert(*models.Product) (*models.Product, error)
}
