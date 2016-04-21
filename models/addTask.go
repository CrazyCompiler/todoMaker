package models

import (
	"database/sql"
	"todoMaker/errorHandler"
)

const dbInsertQuery string = "insert into tasks(task,priority)  VALUES($1,$2) returning taskId;"

func AddTask(db *sql.DB, task string, priority string) error {
	var lastInsertId int
	err := db.QueryRow(dbInsertQuery, task, priority).Scan(&lastInsertId)
	if err != nil {
		errorHandler.DatabaseErrorHandler(err)
		return err
	}
	return nil
}
