//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"fmt"
	"os"
	"io/ioutil"
	"errors"
	"encoding/json"
)

// JsonReadData will open a file and read the data into a struct via json unmarshal
// it returns error if issue occurred
func JsonReadData(filePathName string, data interface{}) (error) {
	//open file
	newFile, err := os.Open(filePathName)
	if err != nil {
		return errors.New(fmt.Sprint("Error opening file:", err))
	}
	defer newFile.Close()

	//read data
	jsonData, err := ioutil.ReadAll(newFile)
	if err != nil {
		return err
	}

	//unmarshal json data to struct
	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return err
	}

	return nil
}

