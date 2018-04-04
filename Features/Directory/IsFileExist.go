// package moviems_directory is a feature package directory which handles anything related to directory functions
// such as create, write, read and delete file as well as check if file exist.
package moviems_directory

import (
	"fmt"
	"os"
)

// Directory.IsFileExist() check if file exist in specified folder
func (d *Directory) IsFileExist(folder, filename string) (bool) {
	fmt.Println("Directory.IsFileExist()")
	if _, err := os.Stat(d.Path + "/" + folder + "/" + filename); os.IsNotExist(err) {
		return false
	}
	return true
}
