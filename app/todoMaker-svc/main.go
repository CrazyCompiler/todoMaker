package main

import (
	"net/http"
	"fmt"
	_"github.com/lib/pq"
	"todoMaker/database"
	"todoMaker/routers"
	//"io/ioutil"
	//"todoMaker/errorHandler"
)

func main()  {
	//dataBaseConfig,err := ioutil.ReadFile("dbConfig")
	//if(err != nil){
	//	errorHandler.DatabaseErrorHandler(err)
	//}
	//
	dbinfo := database.CreateDbInfo();
	db := database.CreateConnection(dbinfo);
	routers.HandleRequests(db)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("their was error ",err)
	}
	defer db.Close()

}
