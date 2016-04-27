package handlers

import (
	"strings"
	"taskManager/models"
	"taskManager/errorHandler"
	"strconv"
	"net/http"
	"taskManager/config"
	"io/ioutil"
)

func AddTask(configObject config.ContextObject) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		taskDescription := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Add(configObject, taskDescription, priority)
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
		err = models.Delete(configObject,task)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UpdateTaskPriority(configObject config.ContextObject)http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		priority := strings.Join(req.Form["priority"], "")
		id,_ := strconv.Atoi(taskId)
		err := models.UpdatePriority(configObject, id,priority)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func UpdateTaskDescription(configObject config.ContextObject)http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		data := strings.Join(req.Form["data"], "")
		id,_ := strconv.Atoi(taskId)
		err := models.UpdateTaskDescription(configObject, id,data)
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}


func UploadCsv(configObject config.ContextObject) http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		file,_,err := req.FormFile("uploadFile")
		if err != nil {
			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
		}
		defer file.Close()
		data,err := ioutil.ReadAll(file)
		err = models.AddTaskByCsv(configObject,string(data))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusOK)
	}
}













//func UploadCsv(configObject config.ContextObject) http.HandlerFunc{
//	return func(res http.ResponseWriter,req *http.Request) {
//		err := req.ParseMultipartForm(32 << 20)
//		if err != nil {
//			errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
//		}
//		m := req.MultipartForm
//
//		files := m.File["uploadFile"]
//
//		for i,_ := range files{
//			file,err := files[i].Open()
//			defer file.Close()
//			if err != nil {
//				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
//			}
//			b1 := make([]byte, 32 << 20)
//			_,err = file.Read(b1)
//			if err != nil {
//				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
//			}
//			separatedData,err := fileReaders.ReadTaskCsv(string(b1))
//			if err != nil {
//				errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
//				fmt.Fprint(res,err.Error())
//				res.WriteHeader(http.StatusBadRequest)
//			}
//
//			for _, each := range separatedData {
//				err := models.Add(configObject,each.TASK ,each.PRIORITY)
//				if err != nil {
//					errorHandler.ErrorHandler(configObject.ErrorLogFile,err)
//					res.WriteHeader(http.StatusInternalServerError)
//					return
//				}
//			}
//			res.WriteHeader(http.StatusCreated)
//		}
//	}
//}

