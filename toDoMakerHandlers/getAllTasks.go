package toDoMakerHandlers

import (
	"net/http"
	"database/sql"
	"todoMaker/errorHandler"
	"encoding/json"
)

type tableContent struct {
	TASKID int
	TASK string
	PRIORITY string
}

func GetAllTasks(db *sql.DB) func(res http.ResponseWriter,req *http.Request){
	return func(res http.ResponseWriter,req *http.Request) {
		dbData := []tableContent{}
		rows, err := db.Query("SELECT * FROM tasks")
		if(err!= nil){
			errorHandler.DatabaseErrorHandler(err)
		}
		for rows.Next(){
			var r tableContent
			err = rows.Scan(&r.TASKID, &r.TASK, &r.PRIORITY)
			if err != nil {
				errorHandler.DatabaseErrorHandler(err)
			}
			dbData = append(dbData,r)
		}
		data,err:= json.Marshal(dbData)
		res.Write(data)
	}
}
