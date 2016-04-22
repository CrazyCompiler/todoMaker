package errorHandler

import (
	"os"
	"time"
	"fmt"
)

func ErrorHandler(errorOccurred error) {
	f, err := os.OpenFile("../errorLog", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	errorBody := time.ANSIC + " : "  +errorOccurred.Error()+ "\n"
	fmt.Print(errorBody)
	if _, err = f.WriteString(errorBody); err != nil {
		panic(err)
	}
}

