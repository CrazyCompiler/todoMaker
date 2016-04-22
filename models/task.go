package models

import (
	"database/sql"
	"taskManager/errorHandler"
	"taskManager/converters"
	"encoding/json"
)

const (
	dbSelectQuery string = "select taskId,task,priority from tasks;"
	dbInsertQuery string = "insert into tasks(task,priority)  VALUES($1,$2) returning taskId;"
	dbDeleteQuery string = "delete from tasks where taskId=$1"
)

func Get(db *sql.DB) []byte {
	rows, err := db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(err)
	}
	dbData := converters.ConvertRowsToStructObjects(rows)
	data, err := json.Marshal(dbData)
	if err != nil {
		errorHandler.ErrorHandler(err)
	}
	return data
}

func Add(db *sql.DB, task string, priority string) error {
	var lastInsertId int
	err := db.QueryRow(dbInsertQuery, task, priority).Scan(&lastInsertId)
	if err != nil {
		errorHandler.ErrorHandler(err)
		return err
	}
	return nil
}

func Delete(db *sql.DB, taskId int) error {
	db.QueryRow(dbDeleteQuery,taskId)
	return nil
}
