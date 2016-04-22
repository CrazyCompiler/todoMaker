package routers

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"taskManager/handlers"
)

func HandleRequests(db *sql.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/uploadCsv",handlers.UploadCsv(db)).Methods("POST")
	r.HandleFunc("/deleteTask/{id:[0-9]+}", handlers.DeleteTask(db)).Methods("DELETE")
	r.HandleFunc("/getAllTasks", handlers.GetTasks(db)).Methods("GET")
	r.HandleFunc("/addTask", handlers.AddTask(db)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
}
