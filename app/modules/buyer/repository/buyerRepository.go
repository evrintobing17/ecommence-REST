package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/buyer"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) buyer.BuyerRepository {
	return &repo{
		db: db,
	}
}

// Add new user to db
func (r *repo) Insert(user models.Buyer) (*models.Buyer, error) {
	db := r.db.Create(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

// Delete existing user
func (r *repo) Delete(userId int) error {
	user := models.Buyer{ID: userId}
	db := r.db.Delete(&user, "id = ?", userId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Get user data by email
func (r *repo) GetByEmail(email string) (*models.Buyer, error) {
	var user models.Buyer

	db := r.db.First(&user, "email = ?", email)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

// Get user data by email
func (r *repo) GetByID(id int) (*models.Buyer, error) {
	var user models.Buyer

	db := r.db.First(&user, "id = ?", id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

func (r *repo) UpdatePartial(updateData map[string]interface{}) (*models.Buyer, error) {
	id := updateData["id"]
	if id == nil {
		return nil, errors.New("field if cannot be empty")
	}
	idString := fmt.Sprintf("%v", id)
	driverID, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	var existingUser models.Buyer
	db := r.db.First(&existingUser, "id=?", driverID)
	if db.Error != nil {
		return nil, db.Error
	}

	db = r.db.Debug().Model(&existingUser).Updates(updateData)
	if db.Error != nil {
		return nil, db.Error
	}

	return &existingUser, nil
}

func (r *repo) ExistByUsername(username string) (bool, error) {
	var user models.Buyer
	if r.db.First(&user, "name=?", username).RecordNotFound() {
		return false, nil
	}
	return true, nil

}

func (r *repo) ExistByEmail(email string) (bool, error) {
	var user models.Buyer
	if r.db.First(&user, "email=?", email).RecordNotFound() {
		return false, nil
	}
	return true, nil
}
