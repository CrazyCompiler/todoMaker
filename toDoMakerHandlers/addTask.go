package toDoMakerHandlers

import (
	"database/sql"
	"net/http"
	"strings"
	"todoMaker/models"
)

func AddTaskHandler(db *sql.DB) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		task := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.AddTask(db, task, priority)
		if err != nil {
			res.WriteHeader(500)
		}
		res.WriteHeader(201)
	}
}
