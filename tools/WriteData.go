//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"os"
	"encoding/json"
)

// WriteData will open a file and write the data into file after data is marshal via json and written to file as string
// it returns error if issue occurred
func WriteData(filePath string, data interface{}) error {
	//open file
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	//marshal data
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//write data to file
	_, err = f.WriteString(string(out))
	if err != nil {
		return err
	}
	return nil
}
