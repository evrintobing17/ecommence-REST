package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	repository := NewProductRepository(db)

	entities := models.Product{
		ID:          1,
		ProductName: "TEST",
		Description: "TEST",
		Price:       1,
		SellerID:    1,
	}
	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.Insert(&entities)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	s.T().Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT(.*)`).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("5"))
		mock.ExpectCommit()
		result, err := repository.Insert(&entities)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

}

func (s *Suite) TestGetByID() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewProductRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetByID(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func (s *Suite) TestGetBySellerID() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewProductRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetBySellerID(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func (s *Suite) TestGetAll() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewProductRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
