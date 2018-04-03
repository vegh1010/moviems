package tools

import (
	"os"
)

/*
 * Create directory at specific path
 */
func CreateDirectory(pathToDirectory string) error {
	if _, err := os.Stat(pathToDirectory); os.IsNotExist(err) {
		os.MkdirAll(pathToDirectory, os.ModePerm)
	} else if err != nil {
		return err
	}

	return nil
}