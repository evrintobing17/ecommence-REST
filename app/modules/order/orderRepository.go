package order

import "github.com/evrintobing17/ecommence-REST/app/models"

type OrderRepository interface {
	GetListOrder(id int) (*[]models.Order, error)
	GetSellerListOrder(id int) (*[]models.Order, error)
	InsertOrder(*models.Order) (*models.Order, error)
	UpdateOrder(updateData map[string]interface{}) (*models.Order, error)
}
