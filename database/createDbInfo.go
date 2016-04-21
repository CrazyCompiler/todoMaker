package database

import "fmt"

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	DB_SCHEMA   = "todoMaker"
)

func CreateDbInfo(dbConfig map[string]string)  string{

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable SEARCH_PATH=%s ",
		dbConfig["DB_USER"],dbConfig["DB_PASSWORD"],dbConfig["DB_NAME"],dbConfig["DB_SCHEMA"])
	return dbinfo
}