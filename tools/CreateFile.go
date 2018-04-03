package tools

import (
	"os"
	"path/filepath"
)

/*
 * Create File at specific path
 */
func CreateFile(pathToDirectory, filename string) (filePath string, err error) {
	// write to file
	if pathToDirectory != "" && filename != "" {
		if _, err = os.Stat(pathToDirectory); os.IsNotExist(err) {
			//create directory if not exists
			err = os.MkdirAll(pathToDirectory, os.ModePerm)
			if err != nil {
				return
			}
		}
		filePath = filepath.Join(pathToDirectory, filename)
	}

	return
}