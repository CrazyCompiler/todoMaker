package modules

import (
	"database/sql"
	"todoMaker/errorHandler"
)

func AddTask(db *sql.DB,task string, priority string) (error){
	var lastInsertId int
	err := db.QueryRow("insert into tasks(task,priority)  VALUES($1,$2) returning taskId;", task, priority).Scan(&lastInsertId)
	if(err != nil) {
		errorHandler.DatabaseErrorHandler(err)
		return err
	}
	return nil
}
