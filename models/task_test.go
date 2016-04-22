package models

import (
	"testing"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)


func TestGet(t *testing.T) {
	db,mock, err := sqlmock.New()
	assert.Nil(t,err)
	rows := sqlmock.NewRows([]string{"hello","High"})
	mock.ExpectQuery("select").WillReturnRows(rows)

	Get(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
