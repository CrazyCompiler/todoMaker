package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"todoMaker/database"
	"todoMaker/fileReaders"
	"todoMaker/routers"
)

func main() {
	configFilePath := "fileReaders/dbConfig.csv"
	if len(os.Args) > 1 {
		configFilePath = os.Args[1]
	}
	dbConfigData := csvReaders.ReadCsv(configFilePath)
	dbinfo := database.CreateDbInfo(dbConfigData)
	db := database.CreateConnection(dbinfo)
	defer db.Close()
	routers.HandleRequests(db)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}
