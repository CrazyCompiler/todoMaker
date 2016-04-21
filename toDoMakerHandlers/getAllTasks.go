package toDoMakerHandlers

import (
	"net/http"
	"database/sql"
	"todoMaker/modules"
)

func GetAllTasksHandler(db *sql.DB) func(res http.ResponseWriter,req *http.Request){
	return func(res http.ResponseWriter,req *http.Request) {
		data := modules.GetAllTasks(db)
		res.Write(data)
	}
}
