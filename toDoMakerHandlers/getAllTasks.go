package toDoMakerHandlers

import (
	"net/http"
	"fmt"
	"database/sql"
	"todoMaker/errorHandler"
	"strconv"
)

func GetAllTasks(db *sql.DB) func(res http.ResponseWriter,req *http.Request){
	return func(res http.ResponseWriter,req *http.Request) {
		rows, err := db.Query("SELECT * FROM tasks")
		taskTable := "<table><tr><td>Task id</td><td> Task Description</td><td> Priority </td></tr>";
		for rows.Next() {
			var taskId int
			var task string
			var priority string
			err = rows.Scan(&taskId, &task, &priority)
			errorHandler.DatabaseErrorHandler(err)
			taskTable += "<tr><td>" + strconv.Itoa(taskId) + "</td><td>" + task + "</td><td>" + priority + "</td></tr>"
		}
		taskTable += "</table>"
		fmt.Fprintf(res, string(taskTable))
	}
}
