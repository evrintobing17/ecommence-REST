package seller

import "github.com/evrintobing17/ecommence-REST/app/models"

type SellerRepository interface {
	Insert(user models.Seller) (*models.Seller, error)
	Delete(userId int) error
	GetByEmail(email string) (*models.Seller, error)
	GetByID(id int) (*models.Seller, error)
	UpdatePartial(updateData map[string]interface{}) (*models.Seller, error)
	ExistByUsername(username string) (bool, error)
	ExistByEmail(email string) (bool, error)
}
