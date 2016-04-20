package database

import (
	"fmt"
	"database/sql"
	"todoMaker/errorHandler"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	DB_SCHEMA   = "todoMaker"
)

func CreateConnection() *sql.DB{
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable SEARCH_PATH=%s ",
		DB_USER, DB_PASSWORD, DB_NAME, DB_SCHEMA)
	db, err := sql.Open("postgres", dbinfo)
	errorHandler.DatabaseErrorHandler(err)
	return db
}

