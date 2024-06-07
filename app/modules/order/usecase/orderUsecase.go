package usecase

import (
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/order"
	"github.com/evrintobing17/ecommence-REST/app/modules/product"
	"github.com/jinzhu/gorm"
)

type orderUC struct {
	orderRepo   order.OrderRepository
	productRepo product.ProductRepository
}

func NewOrderUsecase(orderRepo order.OrderRepository, productRepo product.ProductRepository) order.OrderUsecase {
	return &orderUC{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

// AcceptOrder implements order.OrderUsecase.
func (o *orderUC) AcceptOrder(sellerID, itemid int, alamat string) (*models.Order, error) {
	updateData := map[string]interface{}{
		"id":                           itemid,
		"delivery_destination_address": alamat,
		"status":                       "Accepted",
	}

	isProductExists := true
	products, err := o.productRepo.GetBySellerID(sellerID)
	if err != nil {
		return nil, err
	}
	for _, product := range *products {
		if product.ID != itemid {
			isProductExists = false
		}
	}

	if !isProductExists {
		return nil, gorm.ErrRecordNotFound
	}
	createOrder, err := o.orderRepo.UpdateOrder(updateData)
	if err != nil {
		return nil, err
	}
	return createOrder, nil
}

// CreateOrder implements order.OrderUsecase.
func (o *orderUC) CreateOrder(buyerID, itemID, quantity int, address string) (*models.Order, error) {

	product, err := o.productRepo.GetByID(itemID)
	if err != nil {
		return nil, err
	}
	order := models.Order{
		BuyerID:                    buyerID,
		SellerID:                   product.SellerID,
		DeliveryDestinationAddress: address,
		ItemID:                     itemID,
		Price:                      product.Price,
		Quantity:                   quantity,
		TotalPrice:                 product.Price * float64(quantity),
		Status:                     "Pending",
	}
	creatOrder, err := o.orderRepo.InsertOrder(&order)
	if err != nil {
		return nil, err
	}
	return creatOrder, nil
}

// GetListOrder implements order.OrderUsecase.
func (o *orderUC) GetListOrder(id int) (*[]models.Order, error) {
	order, err := o.orderRepo.GetListOrder(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *orderUC) GetSellerListOrder(id int) (*[]models.Order, error) {
	order, err := o.orderRepo.GetSellerListOrder(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
