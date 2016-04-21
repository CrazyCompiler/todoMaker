package converters

import (
	"database/sql"
	"todoMaker/errorHandler"
)

type tableContent struct {
	TASKID   int
	TASK     string
	PRIORITY string
}

func ConvertRowsToStructObjects(rows *sql.Rows) []tableContent {
	dbData := []tableContent{}
	for rows.Next() {
		var r tableContent
		err := rows.Scan(&r.TASKID, &r.TASK, &r.PRIORITY)
		if err != nil {
			errorHandler.DatabaseErrorHandler(err)
		}
		dbData = append(dbData, r)
	}
	return dbData
}
