// package moviems_directory is a feature package directory which handles anything related to directory functions
// such as create, write, read and delete file as well as check if file exist.
package moviems_directory

import (
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

// Directory.WriteData() write data from data parameter in specified file
func (d *Directory) WriteData(folder, filename string, data interface{}) (err error) {
	fmt.Println("Directory.WriteData()")

	//create file if does not exist
	err = tools.CreateFile(d.Path+"/"+folder, filename)
	if err != nil {
		return
	}

	//write data into file
	err = tools.WriteData(d.Path+"/"+folder+"/"+filename, data)
	if err != nil {
		return
	}
	return
}
