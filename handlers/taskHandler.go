package handlers

import (
	"strings"
	"taskManager/models"
	"taskManager/errorHandler"
	"strconv"
	"taskManager/fileReaders"
	"net/http"
	"taskManager/config"
)

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		task := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Add(configObject, task, priority)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func GetTasks(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data := models.Get(configObject)
		res.WriteHeader(http.StatusOK)
		res.Write(data)
	}
}

func DeleteTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		taskId := strings.Split(req.RequestURI,"/")[2]
		task,err := strconv.Atoi(taskId)
		err = models.Delete(configObject.Db,task)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UploadCsv(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		m := req.MultipartForm

		files := m.File["uploadFile"]

		for i,_ := range files{
			file,err := files[i].Open()
			defer file.Close()
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			b1 := make([]byte, 32 << 20)
			_,err = file.Read(b1)
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			}
			separatedData,err := fileReaders.ReadTaskCsv(string(b1))
			if err != nil {
				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
				res.WriteHeader(http.StatusBadRequest)
			}

			for _, each := range separatedData {
				err := models.Add(configObject,each.TASK ,each.PRIORITY)
				if err != nil {
					errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
					res.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			res.WriteHeader(http.StatusCreated)
		}
	}
}