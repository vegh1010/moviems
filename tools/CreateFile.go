package tools

import (
	"os"
	"path/filepath"
)

/*
 * Create File at specific path
 */
func CreateFile(pathToDirectory, filename string) (filePath string, err error) {
	if pathToDirectory != "" && filename != "" {
		if _, err = os.Stat(pathToDirectory + "/" + filename); os.IsNotExist(err) {
			//create file if not exists
			_, err = os.Create(pathToDirectory+"/"+filename)
			if err != nil {
				return
			}
		}
		filePath = filepath.Join(pathToDirectory, filename)
	}

	return
}
