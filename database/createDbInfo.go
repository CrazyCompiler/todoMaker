package database

import "fmt"

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	DB_SCHEMA   = "todoMaker"
)

func CreateDbInfo()  string{
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable SEARCH_PATH=%s ",
		DB_USER, DB_PASSWORD, DB_NAME, DB_SCHEMA)
	return dbinfo
}
