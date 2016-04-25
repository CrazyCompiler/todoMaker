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

	for i:=0;i<len(rawCSVdata)-1;i++  {
		eachEntry := rawCSVdata[i]
		dataArray = append(dataArray, TableContent{eachEntry[0],eachEntry[1]})
	}
	return dataArray,err

}

