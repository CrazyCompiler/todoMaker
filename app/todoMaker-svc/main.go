package main

import (
	"net/http"
	"fmt"
	_"github.com/lib/pq"
	"todoMaker/database"
	"todoMaker/routers"
	"todoMaker/fileReaders"
)

func main()  {
	dbConfigData := csvReaders.ReadCsv("fileReaders/dbConfig.csv")
	dbinfo := database.CreateDbInfo(dbConfigData);
	db := database.CreateConnection(dbinfo);
	defer db.Close()
	routers.HandleRequests(db)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("their was error ",err)
	}


}
