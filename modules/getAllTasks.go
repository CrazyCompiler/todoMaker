package modules

import (
	"database/sql"
	"todoMaker/errorHandler"
	"encoding/json"
	"todoMaker/converters"
)


func GetAllTasks(db *sql.DB)([]byte){
	rows, err := db.Query("SELECT * FROM tasks")
	if(err!= nil){
		errorHandler.DatabaseErrorHandler(err)
	}
	dbData := converters.ConvertRowsToStructObjects(rows)
	data,err:= json.Marshal(dbData)

	return data
}
