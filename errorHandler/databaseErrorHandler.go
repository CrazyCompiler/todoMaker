package errorHandler

import "io/ioutil"

func DatabaseErrorHandler(err error ){
	if err != nil {
	errorBody := []byte(err.Error())
	ioutil.WriteFile("error.txt",errorBody,0600)
	}
}
