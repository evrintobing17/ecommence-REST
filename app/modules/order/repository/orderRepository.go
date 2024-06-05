package repository

import (
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/order"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) order.OrderRepository {
	return &repo{
		db: db,
	}
}

// GetListOrder implements order.OrderRepository.
func (r *repo) GetListOrder(id int) (*[]models.Order, error) {
	var order []models.Order
	db := r.db.Find(&order, "buyer_id=?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}

// InsertOrder implements order.OrderRepository.
func (r *repo) InsertOrder(order *models.Order) (*models.Order, error) {
	db := r.db.Create(&order)
	if db.Error != nil {
		return nil, db.Error
	}
	return order, nil
}

// UpdateOrder implements order.OrderRepository.
func (r *repo) UpdateOrder(updateData map[string]interface{}) (*models.Order, error) {
	var order models.Order
	db := r.db.Model(&order).Updates(&updateData)
	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}
