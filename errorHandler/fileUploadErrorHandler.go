package errorHandler

import "io/ioutil"

func FileUploadErrorHandler(err error) {
	if err != nil {
		errorBody := []byte(err.Error())
		ioutil.WriteFile("fileUploadError.txt", errorBody, 0600)
	}
}
