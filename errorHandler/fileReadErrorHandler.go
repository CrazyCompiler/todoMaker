package errorHandler

import "io/ioutil"

func FileReadErrorHandler(err error ){
	if err != nil {
		errorBody := []byte(err.Error())
		ioutil.WriteFile("fileReadError.txt",errorBody,0600)
	}
}
