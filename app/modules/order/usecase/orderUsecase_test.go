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

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	productRepo := new(mocks.ProductRepository)

	uc := NewOrderUsecase(mockRepo, productRepo)
	var orderModels models.Order
	gofakeit.Struct(&orderModels)
	entity := &models.Product{
		ID:          1,
		ProductName: "TEST",
		Description: "TEST",
		Price:       0,
		SellerID:    1,
	}

	productRepo.On("GetByID", mock.Anything).Return(entity, nil).Once()
	mockRepo.On("InsertOrder", mock.Anything).Return(&orderModels, nil).Once()

	order, err := uc.CreateOrder(1, 1, 1, "test")

	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestAccOrder_Success(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	productRepo := new(mocks.ProductRepository)

	uc := NewOrderUsecase(mockRepo, productRepo)
	var orderModels models.Order
	gofakeit.Struct(&orderModels)
	var entities []models.Product

	entity := models.Product{
		ID:          1,
		ProductName: "TEST",
		Description: "TEST",
		Price:       1,
		SellerID:    1,
	}

	entities = append(entities, entity, entity)

	productRepo.On("GetBySellerID", mock.Anything).Return(&entities, nil).Once()
	mockRepo.On("UpdateOrder", mock.Anything).Return(&orderModels, nil).Once()

	order, err := uc.AcceptOrder(1, 1, "test")

	assert.Nil(t, err)
	assert.NotNil(t, order)
}

func TestGetListOrder_Error(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	productRepo := new(mocks.ProductRepository)

	uc := NewOrderUsecase(mockRepo, productRepo)

	var entities *[]models.Order
	gofakeit.Struct(&entities)

	mockRepo.On("GetListOrder", mock.Anything).Return(nil, errors.New("failed")).Once()

	product, err := uc.GetListOrder(1)

	assert.NotNil(t, err)
	assert.Nil(t, product)
}

func TestGetListOrder_Success(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	productRepo := new(mocks.ProductRepository)

	uc := NewOrderUsecase(mockRepo, productRepo)

	var entities *[]models.Order
	gofakeit.Struct(&entities)

	mockRepo.On("GetListOrder", mock.Anything).Return(entities, nil).Once()

	product, err := uc.GetListOrder(1)

	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestGetSellerListOrder_Error(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	productRepo := new(mocks.ProductRepository)

	uc := NewOrderUsecase(mockRepo, productRepo)

	var entities *[]models.Order
	gofakeit.Struct(&entities)

	mockRepo.On("GetSellerListOrder", mock.Anything).Return(entities, errors.New("failed")).Once()

	product, err := uc.GetSellerListOrder(1)

	assert.NotNil(t, err)
	assert.Nil(t, product)
}

func TestGetSellerListOrder_Success(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)
	productRepo := new(mocks.ProductRepository)

	uc := NewOrderUsecase(mockRepo, productRepo)

	var entities *[]models.Order
	gofakeit.Struct(&entities)

	mockRepo.On("GetSellerListOrder", mock.Anything).Return(entities, nil).Once()

	product, err := uc.GetSellerListOrder(1)

	assert.Nil(t, err)
	assert.NotNil(t, product)
}
