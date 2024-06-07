package models

type Order struct {
	ID                         int     `json:"id"`
	BuyerID                    int     `json:"buyer_id"`
	SellerID                   int     `json:"seller_id"`
	DeliverySourceAddress      string  `json:"delivery_source_address"`
	DeliveryDestinationAddress string  `json:"delivery_destination_address"`
	ItemID                     int     `json:"item_id"`
	Price                      float64 `json:"price"`
	Quantity                   int     `json:"quantity"`
	TotalPrice                 float64 `json:"total_price"`
	Status                     string  `json:"status"`
}

type OrderRequest struct {
	ItemID   int `json:"item_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required,gt=0"`
}

type AcceptOrderRequest struct {
	ItemID   int `json:"item_id" binding:"required"`
}

func (o Order) TableName() string {
	return "ecommerce.order"
}
