package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"taskManager/database"
	"taskManager/fileReaders"
	"taskManager/routers"
)

func main() {
	configFilePath := "fileReaders/demoJson"
	if len(os.Args) > 1 {
		configFilePath = os.Args[1]
	}
	dbConfigDataJson := fileReaders.ReadJsonFile(configFilePath)
	dbinfo := database.CreateDbInfo(dbConfigDataJson)

	db := database.CreateConnection(dbinfo)
	defer db.Close()
	routers.HandleRequests(db)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}
