package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	helpers "github.com/evrintobing17/ecommence-REST/app/helpers/sqlMock"
	"github.com/evrintobing17/ecommence-REST/app/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
}

func TestSqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestCreate() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewOrderRepository(db)

	entities := models.Order{
		ID:                         1,
		BuyerID:                    1,
		SellerID:                   1,
		DeliverySourceAddress:      "TEST",
		DeliveryDestinationAddress: "TEST",
		ItemID:                     1,
		Price:                      1,
		Quantity:                   1,
		TotalPrice:                 1,
		Status:                     "TEST",
	}
	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.InsertOrder(&entities)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	s.T().Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT(.*)`).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("5"))
		mock.ExpectCommit()
		result, err := repository.InsertOrder(&entities)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

}

func (s *Suite) TestGetListOrder() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewOrderRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetListOrder(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func (s *Suite) TestGetSellerListOrder() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewOrderRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetSellerListOrder(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func (s *Suite) TestUpdateOrder() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)
	var entities models.Seller
	if err := gofakeit.Struct(&entities); err != nil {
		s.T().Fatal(err)
	}
	repository := NewOrderRepository(db)

	s.T().Run("success", func(t *testing.T) {
		query := `SELECT * FROM "ecommerce"."seller" WHERE (id=$1) ORDER BY "ecommerce"."seller"."id" ASC LIMIT 1`
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(entities.ID, entities.Name))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE(.*)`).WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		result, err := repository.UpdateOrder(map[string]interface{}{"id": "1"})
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
