package fileReaders

import (
	"encoding/csv"
	"strings"
	"taskManager/errorHandler"
)

const totalNoOfRows = 2

func ReadTaskCsv(fileData string) ([]map[string]string,error) {
	dataArray := []map[string]string{}
	reader := csv.NewReader(strings.NewReader(fileData))

	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		errorHandler.ErrorHandler(err)
	}

	for _, each := range rawCSVdata {
		entry := make(map[string]string)
		if(len(each) == totalNoOfRows) {
			entry["task"] = each[0]
			entry["priority"] = each[1]
			dataArray = append(dataArray, entry)
		}
	}
	return dataArray,err
}

