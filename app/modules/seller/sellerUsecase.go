package seller

import (
	"github.com/evrintobing17/ecommence-REST/app/models"
)

type SellerUsecase interface {
	Login(email, password string) (user *models.Seller, token string, err error)
	Register(username, email, password, age string, isbuyer bool) (user *models.Seller, err error)
	RefreshAccessJWT(userID int) (newAccessJWT string, err error)
	DeleteByID(userId int) error
	Update(updateData map[string]interface{}) (user *models.Seller, err error)
}
