package routers

import (
	"database/sql"
	"net/http"
	"todoMaker/toDoMakerHandlers"
)

func staticFiles(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}

func HandleRequests(db *sql.DB) {
	http.HandleFunc("/getAllTasks", toDoMakerHandlers.GetAllTasksHandler(db))
	http.HandleFunc("/addTask", toDoMakerHandlers.AddTaskHandler(db))
	http.HandleFunc("/", staticFiles)
}
