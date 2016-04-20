package toDoMakerHandlers

import (
	"net/http"
	"strings"
	"database/sql"
	"todoMaker/errorHandler"
)

func AddTask(db *sql.DB) func(res http.ResponseWriter,req *http.Request){
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm();
		task := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		var lastInsertId int
		err := db.QueryRow("insert into tasks(task,priority)  VALUES($1,$2) returning taskId;", task, priority).Scan(&lastInsertId)
		errorHandler.DatabaseErrorHandler(err)

		GetAllTasks(db)(res,req);
	}
}


