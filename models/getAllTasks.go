package models

import (
	"database/sql"
	"encoding/json"
	"todoMaker/converters"
	"todoMaker/errorHandler"
)

const dbQuery string = "SELECT * FROM tasks"

func GetAllTasks(db *sql.DB) []byte {
	rows, err := db.Query(dbQuery)
	if err != nil {
		errorHandler.DatabaseErrorHandler(err)
	}
	dbData := converters.ConvertRowsToStructObjects(rows)
	data, err := json.Marshal(dbData)

	return data
}
