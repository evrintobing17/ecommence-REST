package repository

import (
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/product"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	return &repo{
		db: db,
	}
}

// GetAll implements product.ProductRepository.
func (r *repo) GetAll() (*[]models.Product, error) {
	var product []models.Product
	db := r.db.Find(&product)
	if db.Error != nil {
		return nil, db.Error
	}
	return &product, nil
}

// GetByID implements product.ProductRepository.
func (r *repo) GetBySellerID(id int) (*[]models.Product, error) {
	var product []models.Product
	db := r.db.Find(&product, "seller_id=?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &product, nil
}

// Insert implements product.ProductRepository.
func (r *repo) Insert(product *models.Product) (*models.Product, error) {
	db := r.db.Create(&product)
	if db.Error != nil {
		return nil, db.Error
	}
	return product, nil
}

func (r *repo) GetByID(id int) (*models.Product, error) {
	var product models.Product
	db := r.db.First(&product, "seller_id=?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &product, nil
}
