package moviems_directory

import (
	"fmt"
	"os"
)

func (d *Directory) IsFileExist(folder, filename string) (bool) {
	fmt.Println("Directory.IsFileExist()")
	if _, err := os.Stat(d.Path + "/" + folder + "/" + filename); os.IsNotExist(err) {
		return false
	}
	return true
}
