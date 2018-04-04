// package moviems_directory is a feature package directory which handles anything related to directory functions
// such as create, write, read and delete file as well as check if file exist.
package moviems_directory

import (
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

// Directory.CreateFile() create a file in the specified folder
func (d *Directory) CreateFile(folder, filename string) (err error) {
	fmt.Println("Directory.CreateFile()")
	err = tools.CreateFile(d.Path+"/"+folder, filename)
	if err != nil {
		return
	}
	return
}
