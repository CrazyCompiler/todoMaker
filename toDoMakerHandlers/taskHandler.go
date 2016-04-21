package toDoMakerHandlers

import (
	"database/sql"
	"strings"
	"todoMaker/models"
	"net/http"
	"strconv"
	"fmt"
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
		req.ParseMultipartForm(32 << 20)
		file,handler,err := req.FormFile("uploadFile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		file,err := handler.Open()
		fmt.Println(file,err)


	}
}