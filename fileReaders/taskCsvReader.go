package fileReaders

import (
	"encoding/csv"
	"strings"
	"taskManager/validate"
)

type TableContent struct {
	TASK     string
	PRIORITY string
}



func ReadTaskCsv(fileData string) ([]TableContent,error) {
	dataArray := []TableContent{}
	reader := csv.NewReader(strings.NewReader(fileData))

	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	err = validate.ValidateAllEntry(rawCSVdata)
	if err!=nil {
		return dataArray,err
	}

	for _, each := range rawCSVdata {
		entry := TableContent{}
		if(len(each) == 2) {
			entry.TASK = each[0]
			entry.PRIORITY = each[1]
			dataArray = append(dataArray, entry)
		}
	}
	return dataArray,err

}

