package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskManager/handlers"
	"taskManager/config"
)


func HandleRequests(configObject config.ContextObject) {
	r := mux.NewRouter()
	r.HandleFunc("/updateTaskDescription",handlers.UpdateTaskDescription(configObject)).Methods("POST")
	r.HandleFunc("/updatePriority",handlers.UpdateTaskPriority(configObject)).Methods("POST")
	r.HandleFunc("/uploadCsv",handlers.UploadCsv(configObject)).Methods("POST")
	r.HandleFunc("/deleteTask/{id:[0-9]+}", handlers.DeleteTask(configObject)).Methods("DELETE")
	r.HandleFunc("/getAllTasks", handlers.GetTasks(configObject)).Methods("GET")
	r.HandleFunc("/addTask", handlers.AddTask(configObject)).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
}
