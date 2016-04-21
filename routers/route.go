package routers

import (
	"database/sql"
	"net/http"
	"todoMaker/toDoMakerHandlers"
	"github.com/gorilla/mux"
)

func HandleRequests(db *sql.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/getAllTasks", toDoMakerHandlers.GetAllTasksHandler(db))
	r.HandleFunc("/addTask", toDoMakerHandlers.AddTaskHandler(db))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)

}
