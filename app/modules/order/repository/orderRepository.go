package repository

import (
	"fmt"

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

func (r *repo) GetSellerListOrder(id int) (*[]models.Order, error) {
	var order []models.Order
	db := r.db.Find(&order, "seller_id=?", id)
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
	fmt.Println(updateData, "apa")
	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}
