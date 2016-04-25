package models

import (
	"taskManager/errorHandler"
	"taskManager/converters"
	"encoding/json"
	"taskManager/config"
)

const (
	dbSelectQuery string = "select taskId,task,priority from tasks;"
	dbInsertQuery string = "insert into tasks(task,priority)  VALUES($1,$2) returning taskId;"
	dbDeleteQuery string = "delete from tasks where taskId=$1"
	dbPriorityUpdateQuery string = "update tasks set priority=$1 where taskID=$2;"
)

func Get(configObject config.ContextObject) []byte {
	rows, err := configObject.Db.Query(dbSelectQuery)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	dbData := converters.ConvertRowsToStructObjects(rows)
	data, err := json.Marshal(dbData)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
	}
	return data
}

func Add(configObject config.ContextObject, taskDescription string, priority string) error {
	task := Task{taskDescription,priority}
	return task.Create(configObject)
}

func Delete(configObject config.ContextObject, taskId int) error {
	_,err := configObject.Db.Exec(dbDeleteQuery,taskId)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePriority(configObject config.ContextObject, taskId int, priority string)error{
	_,err := configObject.Db.Exec(dbPriorityUpdateQuery, priority, taskId)
	if err != nil {
		errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		return err
	}
	return nil
}