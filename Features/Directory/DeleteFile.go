// package moviems_directory is a feature package directory which handles anything related to directory functions
// such as create, write, read and delete file as well as check if file exist.
package moviems_directory

import (
	"fmt"
	"os"
)

// Directory.DeleteFile() delete a file in the specified folder
func (d *Directory) DeleteFile(folder, filename string) (err error) {
	fmt.Println("Directory.DeleteFile()")
	err = os.Remove(d.Path + "/" + folder + "/" + filename)
	if err != nil {
		return
	}
	return
}
