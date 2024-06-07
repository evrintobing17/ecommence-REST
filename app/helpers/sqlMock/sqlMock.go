package helpers

import (
	"database/sql/driver"
	"reflect"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func NewMockDB(typesConn string) (*gorm.DB, sqlmock.Sqlmock, error) {

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open("postgres", sqlDB)

	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func GetValueAndColumnStructToDriverValue(value interface{}) ([]driver.Value, []string) {
	var result []driver.Value

	object := reflect.ValueOf(value)
	var column []string
	for i := 0; i < object.NumField(); i++ {

		// set data to driver.Value
		result = append(result, object.Field(i).Interface())

		// set column
		splitStringByTag := strings.Split(object.Type().Field(i).Tag.Get("gorm"), ";")
		for i := 0; i < len(splitStringByTag); i++ {
			if strings.Contains(splitStringByTag[i], "column") {
				column = append(column, strings.TrimPrefix(splitStringByTag[i], "column:"))
				break
			}
		}
	}

	return result, column
}
