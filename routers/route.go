package routers

import (
	"database/sql"
	"net/http"
	"todoMaker/toDoMakerHandlers"
	"github.com/gorilla/mux"
)

func HandleRequests(db *sql.DB) {
	r := mux.NewRouter()
	r.HandleFunc("/uploadCsv",toDoMakerHandlers.UploadCsv(db)).Methods("POST")
	r.HandleFunc("/deleteTask", toDoMakerHandlers.DeleteTask(db)).Methods("POST")
	r.HandleFunc("/getAllTasks", toDoMakerHandlers.GetTasks(db)).Methods("GET")
	r.HandleFunc("/addTask", toDoMakerHandlers.AddTask(db)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
}
