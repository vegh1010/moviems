//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"os"
)

// CreateFile will verify if file does not exist then create file at specified path
// it returns an error if issue occurred
func CreateFile(pathToDirectory, filename string) (err error) {
	//check if file exist
	if _, err = os.Stat(pathToDirectory + "/" + filename); os.IsNotExist(err) {
		//create file
		_, err = os.Create(pathToDirectory+"/"+filename)
		if err != nil {
			return
		}
	}

	return
}
