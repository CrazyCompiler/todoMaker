package routers

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"taskManager/handlers"
)

func HandleRequests(db *sql.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/uploadCsv",toDoMakerHandlers.UploadCsv(db)).Methods("POST")
	r.HandleFunc("/deleteTask/{id:[0-9]+}", toDoMakerHandlers.DeleteTask(db)).Methods("DELETE")
	r.HandleFunc("/getAllTasks", toDoMakerHandlers.GetTasks(db)).Methods("GET")
	r.HandleFunc("/addTask", toDoMakerHandlers.AddTask(db)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
}
