package main

import (
	"net/http"
	"fmt"
	_"github.com/lib/pq"
	"todoMaker/toDoMakerHandlers"
	"todoMaker/database"
)

func staticFiles(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}


func main()  {
	db := database.CreateConnection();
	http.HandleFunc("/getAllTasks",toDoMakerHandlers.GetAllTasks(db))
	http.HandleFunc("/addTask", toDoMakerHandlers.AddTask(db))
	http.HandleFunc("/", staticFiles)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("their was error ",err)
	}
	defer db.Close()
}
