package buyer

import (
	"github.com/evrintobing17/ecommence-REST/app/models"
)

type BuyerUsecase interface {
	Login(email, password string) (user *models.Buyer, token string, err error)
	Register(username, email, password, address string) (user *models.Buyer, err error)
	RefreshAccessJWT(userID int) (newAccessJWT string, err error)
	DeleteByID(userId int) error
	Update(updateData map[string]interface{}) (user *models.Buyer, err error)
}
