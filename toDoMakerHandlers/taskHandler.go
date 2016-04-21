package toDoMakerHandlers

import (
	"database/sql"
	"strings"
	"todoMaker/models"
	"net/http"
	"strconv"
	"todoMaker/errorHandler"
	"todoMaker/fileReaders"
)

func AddTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		task := strings.Join(req.Form["task"], "")
		priority := strings.Join(req.Form["priority"], "")
		err := models.Add(db, task, priority)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusCreated)
	}
}

func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data := models.Get(db)
		res.WriteHeader(http.StatusOK)
		res.Write(data)
	}
}

func DeleteTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter,req *http.Request) {
		req.ParseForm()
		taskId := strings.Join(req.Form["taskId"], "")
		task,err := strconv.Atoi(taskId)
		err = models.Delete(db,task)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.WriteHeader(http.StatusAccepted)
	}
}

func UploadCsv(db *sql.DB) http.HandlerFunc{
	return func(res http.ResponseWriter,req *http.Request) {
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			errorHandler.FileUploadErrorHandler(err)
		}
		m := req.MultipartForm
		files := m.File["uploadFile"]
		for i,_ := range files{
			file,err := files[i].Open()
			defer file.Close()
			if err != nil {
				errorHandler.FileUploadErrorHandler(err)
			}
			b1 := make([]byte, 32 << 20)
			_,err = file.Read(b1)
			seperatedData := fileReaders.ReadTaskCsv(string(b1))
			for _, each := range seperatedData {
				err := models.Add(db,each["task"],each["priority"])
				if err != nil {
					res.WriteHeader(http.StatusInternalServerError)
				}
				res.WriteHeader(http.StatusCreated)
			}
		}
	}
}