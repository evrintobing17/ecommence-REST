package buyer

import "github.com/evrintobing17/ecommence-REST/app/models"

type BuyerRepository interface {
	//buyer
	Insert(user models.Buyer) (*models.Buyer, error)
	Delete(userId int) error
	GetByEmail(email string) (*models.Buyer, error)
	GetByID(id int) (*models.Buyer, error)
	UpdatePartial(updateData map[string]interface{}) (*models.Buyer, error)
	ExistByUsername(username string) (bool, error)
	ExistByEmail(email string) (bool, error)
}
