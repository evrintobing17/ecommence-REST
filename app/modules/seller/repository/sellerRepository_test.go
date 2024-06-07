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

	repository := NewSellerRepository(db)

	entities := models.Seller{
		ID:       1,
		Name:     "TEST",
		Email:    "TEST",
		Password: "TEST",
		Address:  "TEST",
	}
	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.Insert(entities)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	s.T().Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT(.*)`).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("5"))
		mock.ExpectCommit()
		result, err := repository.Insert(entities)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

}

func (s *Suite) TestGetByID() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewSellerRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetByID(1)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func (s *Suite) TestGetByEmail() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)

	repository := NewSellerRepository(db)

	s.T().Run("error create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`SELECT(.*)`).WithArgs().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		result, err := repository.GetByEmail("")
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func (s *Suite) TestUpdate() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)
	var entities models.Seller
	if err := gofakeit.Struct(&entities); err != nil {
		s.T().Fatal(err)
	}
	repository := NewSellerRepository(db)

	s.T().Run("error id nil", func(t *testing.T) {
		result, err := repository.UpdatePartial(map[string]interface{}{"id": ""})
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	s.T().Run("success", func(t *testing.T) {
		query := `SELECT * FROM "ecommerce"."seller" WHERE (id=$1) ORDER BY "ecommerce"."seller"."id" ASC LIMIT 1`
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(entities.ID, entities.Name))
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE(.*)`).WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		result, err := repository.UpdatePartial(map[string]interface{}{"id": "1"})
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func (s *Suite) TestGetByName() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)
	var entities models.Seller
	if err := gofakeit.Struct(&entities); err != nil {
		s.T().Fatal(err)
	}
	repository := NewSellerRepository(db)

	s.T().Run("error update", func(t *testing.T) {
		query := `SELECT * FROM "ecommerce"."seller" WHERE (name=$1) ORDER BY "ecommerce"."seller"."id" ASC LIMIT 1`
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(entities.ID, entities.Name))
		result, err := repository.ExistByUsername("TEST")
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func (s *Suite) TestExistEmail() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)
	var entities models.Seller
	if err := gofakeit.Struct(&entities); err != nil {
		s.T().Fatal(err)
	}
	repository := NewSellerRepository(db)

	s.T().Run("error update", func(t *testing.T) {
		query := `SELECT * FROM "ecommerce"."seller" WHERE (email=$1) ORDER BY "ecommerce"."seller"."id" ASC LIMIT 1`
		queryRegex := regexp.QuoteMeta(query)
		mock.ExpectQuery(queryRegex).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(entities.ID, entities.Name))
		result, err := repository.ExistByEmail("TEST")
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func (s *Suite) TestDelete() {
	db, mock, err := helpers.NewMockDB("postgres")
	require.NoError(s.T(), err)
	var entities models.Seller
	if err := gofakeit.Struct(&entities); err != nil {
		s.T().Fatal(err)
	}
	repository := NewSellerRepository(db)

	s.T().Run("error update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE(.*)").WithArgs().WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		err := repository.Delete(1)
		assert.Nil(t, err)
	})
}
