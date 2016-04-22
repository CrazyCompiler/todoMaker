package fileReaders

import (
	"encoding/csv"
	"strings"
)

type TableContent struct {
	TASK     string
	PRIORITY string
}


const totalNoOfRows = 2

func ReadTaskCsv(fileData string) ([]TableContent,error) {
	dataArray := []TableContent{}
	reader := csv.NewReader(strings.NewReader(fileData))

	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	for _, each := range rawCSVdata {
		entry := TableContent{}
		if(len(each) == totalNoOfRows) {
			entry.TASK = each[0]
			entry.PRIORITY = each[1]
			dataArray = append(dataArray, entry)
		}
	}
	return dataArray,err

}

