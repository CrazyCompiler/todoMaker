package converters

import (
	"database/sql"
	"taskManager/errorHandler"
)

type TableContent struct {
	TASKID   int
	TASK     string
	PRIORITY string
}

func ConvertRowsToStructObjects(rows *sql.Rows) []TableContent {
	dbData := []TableContent{}
	for rows.Next() {
		var r TableContent
		err := rows.Scan(&r.TASKID, &r.TASK, &r.PRIORITY)
		if err != nil {
			errorHandler.ErrorHandler(err)
		}
		dbData = append(dbData, r)
	}
	return dbData
}
