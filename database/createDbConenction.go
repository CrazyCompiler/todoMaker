package database

import (
	"database/sql"
	"taskManager/errorHandler"
)

func CreateConnection(dbinfo string) *sql.DB {
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		errorHandler.ErrorHandler(err)
	}
	db.Ping()
	return db
}
