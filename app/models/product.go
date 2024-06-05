package models

type Product struct {
	ID          int     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerID    int     `json:"seller_id"`
}

type Products struct {
	ID          int     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (p Product) TableName() string {
	return "ecommerce.product"
}
