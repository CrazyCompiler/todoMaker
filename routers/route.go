package routers

import (
	"net/http"
	"todoMaker/toDoMakerHandlers"
	"database/sql"
)


func staticFiles(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}


func HandleRequests(db *sql.DB){
	http.HandleFunc("/getAllTasks",toDoMakerHandlers.GetAllTasks(db))
	http.HandleFunc("/addTask", toDoMakerHandlers.AddTask(db))
	http.HandleFunc("/", staticFiles)
}
