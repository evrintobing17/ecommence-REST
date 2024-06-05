package usecase

import (
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/product"
)

type productUC struct {
	productRepo product.ProductRepository
}

func NewProductUsecase(productRepo product.ProductRepository) product.ProductUsecase {
	return &productUC{
		productRepo: productRepo,
	}
}

// CreateProduct implements product.ProductUsecase.
func (p *productUC) CreateProduct(product models.Product) (*models.Product, error) {
	products, err := p.productRepo.Insert(&product)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetALL implements product.ProductUsecase.
func (p *productUC) GetAll() (*[]models.Product, error) {
	product, err := p.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetBySellerID implements product.ProductUsecase.
func (p *productUC) GetBySellerID(id int) (*[]models.Product, error) {
	product, err := p.productRepo.GetBySellerID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
