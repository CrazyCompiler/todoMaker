package csvReaders

import (
	"os"
	"encoding/csv"
	"todoMaker/errorHandler"
)

func ReadCsv(fileName string) map[string]string {
	csvObject := make(map[string]string)
	csvfile, err := os.Open(fileName)

	if err != nil {
		errorHandler.FileReadErrorHandler(err)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		errorHandler.FileReadErrorHandler(err)
	}

	for _, each := range rawCSVdata {
		csvObject[each[0]] = each[1]
	}

	return csvObject
}
