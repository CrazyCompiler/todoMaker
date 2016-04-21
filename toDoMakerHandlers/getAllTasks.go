package toDoMakerHandlers

import (
	"database/sql"
	"net/http"
	"todoMaker/models"
)

func GetAllTasksHandler(db *sql.DB) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		data := models.GetAllTasks(db)
		res.Write(data)
	}
}
