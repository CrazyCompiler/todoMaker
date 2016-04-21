package fileReaders

import (
	"todoMaker/errorHandler"
	"encoding/csv"
	"strings"
)

func ReadTaskCsv(fileData string) []map[string]string {
	 dataArray := []map[string]string{}
	reader := csv.NewReader(strings.NewReader(fileData))

	reader.FieldsPerRecord = -1


	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		errorHandler.FileReadErrorHandler(err)
	}


	for _, each := range rawCSVdata {
		entry := make(map[string]string)
		entry["task"] = each[0]
		entry["priority"] = each[1]
		dataArray = append(dataArray, entry)
	}
	return dataArray
}

