package order

import "github.com/evrintobing17/ecommence-REST/app/models"

type OrderUsecase interface {
	GetListOrder(int) (*[]models.Order, error)
	CreateOrder(buyerID, itemID, quantity int, address string) (*models.Order, error)
	AcceptOrder(itemID int, address string) (*models.Order, error)
}
