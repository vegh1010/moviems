package tools

import (
	"os"
	"encoding/json"
)

func WriteData(filePath string, data interface{}) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(out))
	if err != nil {
		return err
	}
	return nil
}
