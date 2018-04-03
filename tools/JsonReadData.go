package tools

import (
	"fmt"
	"os"
	"io/ioutil"
	"errors"
	"encoding/json"
)

func JsonReadData(filePathName string, data interface{}) (error) {
	newFile, err := os.Open(filePathName)
	if err != nil {
		return errors.New(fmt.Sprint("Error opening file:", err))
	}
	defer newFile.Close()

	jsonData, err := ioutil.ReadAll(newFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return err
	}

	return nil
}

