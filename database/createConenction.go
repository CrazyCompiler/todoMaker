package database

import (
	"database/sql"
	"todoMaker/errorHandler"
)

func CreateConnection(dbinfo string) *sql.DB{
	db, err := sql.Open("postgres", dbinfo)
	errorHandler.DatabaseErrorHandler(err)
	return db
}

