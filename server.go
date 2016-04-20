package main

import (
	"net/http"
	"fmt"
	_"github.com/lib/pq"
	"todoMaker/database"
	"todoMaker/routers"
)

func main()  {
	db := database.CreateConnection();
	routers.HandleRequests(db)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("their was error ",err)
	}
	defer db.Close()

}
