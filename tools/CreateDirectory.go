//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"os"
)

// CreateDirectory will verify if directory does not exist then create directory at specified path
// it returns an error if issue occurred
func CreateDirectory(pathToDirectory string) error {
	if _, err := os.Stat(pathToDirectory); os.IsNotExist(err) {
		os.MkdirAll(pathToDirectory, os.ModePerm)
	} else if err != nil {
		return err
	}

	return nil
}