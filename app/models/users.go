package models

type Buyer struct {
	ID       int    `gorm:"primaryKey;autoIncrement:true" json:"id" example:"1"`
	Name     string `json:"name" example:"Nanda" binding:"required,string"`
	Email    string `json:"email" example:"your.email@example.com"`
	Password string `json:"password" example:"qwerty123"`
	Address  string `json:"address" gorm:"column:alamat_pengiriman"  example:"jl.suka maju"`
	// DateAudit
}

type BuyerQuery struct {
	Name     string `json:"name" example:"Nanda" binding:"required,string"`
	Email    string `json:"email" example:"your.email@example.com"`
	Password string `json:"password" example:"qwerty123"`
	Address  string `json:"address" gorm:"column:alamat_pengiriman"  example:"jl.suka maju"`
}

type Seller struct {
	ID       int    `gorm:"primaryKey;autoIncrement:true" json:"id" example:"1"`
	Name     string `json:"name" example:"Nanda" binding:"required,string"`
	Email    string `json:"email" example:"your.email@example.com"`
	Password string `json:"password" example:"qwerty123"`
	Address  string `json:"address" gorm:"column:alamat_pickup" example:"jl.suka maju"`
}
type SellerQuery struct {
	Name     string `json:"name" example:"Nanda" binding:"required,string"`
	Email    string `json:"email" example:"your.email@example.com"`
	Password string `json:"password" example:"qwerty123"`
	Address  string `json:"address" gorm:"column:alamat_pickup" example:"jl.suka maju"`
}

func (u Buyer) TableName() string {
	return "ecommerce.buyer"
}

func (u Seller) TableName() string {
	return "ecommerce.seller"
}
