package usecase

import (
	"errors"
	"os"
	"testing"

	"github.com/evrintobing17/ecommence-REST/app/mocks"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/evrintobing17/ecommence-REST/app/modules/buyer/usecase"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_NonExistentEmail(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	email := "nonexistent@example.com"
	password := "TESTS"

	mockRepo.On("GetByEmail", email).Return(nil, gorm.ErrRecordNotFound).Once()

	returnedUser, token, err := uc.Login(email, password)

	assert.Error(t, err)
	assert.Equal(t, usecase.ErrInvalidCredential, err)
	assert.Empty(t, token)
	assert.Nil(t, returnedUser)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	email := "nonexistent@example.com"
	password := "TEST123"
	os.Setenv("jwt.expirationDurationDay", "1")
	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "$2a$10$rTPNBWbPXLH1MxqUh11xaOn8IwVgT5QVXE8wvhNvcAeRDcNHRVrJK",
		Address:  "TEST",
	}
	mockRepo.On("GetByEmail", email).Return(&entities, nil).Once()

	returnedUser, token, err := uc.Login(email, password)

	assert.NotNil(t, err)
	assert.NotNil(t, token)
	assert.Nil(t, returnedUser)
}

func TestLogin_SuccessLogin(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	email := "nonexistent@example.com"
	password := "TEST"
	os.Setenv("jwt.expirationDurationDay", "1")
	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "$2a$10$rTPNBWbPXLH1MxqUh11xaOn8IwVgT5QVXE8wvhNvcAeRDcNHRVrJK",
		Address:  "TEST",
	}
	mockRepo.On("GetByEmail", email).Return(&entities, nil).Once()

	returnedUser, token, err := uc.Login(email, password)

	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.NotNil(t, returnedUser)
}

func TestRegister_ExistByUsername(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	os.Setenv("jwt.expirationDurationDay", "1")
	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "TEST",
		Address:  "TEST",
	}
	mockRepo.On("ExistByUsername", entities.Name).Return(true, errors.New("username already exist")).Once()

	returnedUser, err := uc.Register(entities.Name, entities.Email, entities.Password, entities.Address)

	assert.NotNil(t, err)
	assert.Nil(t, returnedUser)
}

func TestRegister_ExistByEmail(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	os.Setenv("jwt.expirationDurationDay", "1")
	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "TEST",
		Address:  "TEST",
	}
	mockRepo.On("ExistByUsername", entities.Name).Return(false, nil).Once()
	mockRepo.On("ExistByEmail", entities.Email).Return(true, errors.New("email already exist")).Once()

	returnedUser, err := uc.Register(entities.Name, entities.Email, entities.Password, entities.Address)

	assert.NotNil(t, err)
	assert.Nil(t, returnedUser)
}

func TestRegister_SuccessRegister(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	os.Setenv("jwt.expirationDurationDay", "1")
	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "TEST",
		Address:  "TEST",
	}
	mockRepo.On("ExistByUsername", entities.Name).Return(false, nil).Once()
	mockRepo.On("ExistByEmail", entities.Email).Return(false, nil).Once()
	mockRepo.On("Insert", mock.Anything).Return(&entities, nil).Once()

	returnedUser, err := uc.Register(entities.Name, entities.Email, entities.Password, entities.Address)

	assert.Nil(t, err)
	assert.NotNil(t, returnedUser)
}

func TestGenerateJWT_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	jwt, err := uc.RefreshAccessJWT(1)

	assert.Nil(t, err)
	assert.NotNil(t, jwt)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)
	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "TEST",
		Address:  "TEST",
	}
	mockRepo.On("UpdatePartial", mock.Anything).Return(&entities, nil).Once()

	user, err := uc.Update(map[string]interface{}{"id": "1"})

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestDelete_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepository)
	uc := NewSellerUsecase(mockRepo)

	mockRepo.On("Delete", mock.Anything).Return(nil).Once()

	err := uc.DeleteByID(1)

	assert.Nil(t, err)
}
