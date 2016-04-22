package errorHandler

import (
	"os"
	"time"
	"fmt"
)

func ErrorHandler(errorFile *os.File,errorOccurred error) {
	errorBody := time.ANSIC + " : "  +errorOccurred.Error()+ "\n"
	fmt.Print(errorBody)
	if _, err := errorFile.WriteString(errorBody); err != nil {
		panic(err)
	}
}

