// package moviems_directory is a feature package directory which handles anything related to directory functions
// such as create, write, read and delete file as well as check if file exist.
package moviems_directory

import (
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

// Directory.ReadData() read data in specified file and add the information into data
func (d *Directory) ReadData(folder, filename string, data interface{}) (err error) {
	fmt.Println("Directory.ReadData()")
	err = tools.JsonReadData(d.Path + "/" + folder+"/"+filename, data)
	if err != nil {
		return
	}
	return
}


