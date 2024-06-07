package usecase

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/evrintobing17/ecommence-REST/app/mocks"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct_Success(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	uc := NewProductUsecase(mockRepo)

	entity := &models.Product{
		ID:          1,
		ProductName: "TEST",
		Description: "TEST",
		Price:       0,
		SellerID:    1,
	}

	mockRepo.On("Insert", mock.Anything).Return(entity, nil).Once()

	product, err := uc.CreateProduct(*entity)

	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestGetAll_Success(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	uc := NewProductUsecase(mockRepo)

	var entities *[]models.Product
	gofakeit.Struct(&entities)
	mockRepo.On("GetAll", mock.Anything).Return(entities, nil).Once()

	product, err := uc.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, product)
}
func TestGetAll_Error(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	uc := NewProductUsecase(mockRepo)

	var entities *[]models.Product
	gofakeit.Struct(&entities)
	mockRepo.On("GetAll", mock.Anything).Return(nil, errors.New("failed")).Once()

	product, err := uc.GetAll()

	assert.NotNil(t, err)
	assert.Nil(t, product)
}

func TestGetBySellerID_Success(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)
	uc := NewProductUsecase(mockRepo)

	var entities *[]models.Product
	gofakeit.Struct(&entities)

	mockRepo.On("GetBySellerID", mock.Anything).Return(entities, nil).Once()

	product, err := uc.GetBySellerID(1)

	assert.Nil(t, err)
	assert.NotNil(t, product)
}
