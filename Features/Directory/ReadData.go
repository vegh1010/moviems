package moviems_directory

import (
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

//write data in specified file
func (d *Directory) ReadData(folder, filename string, data interface{}) (err error) {
	fmt.Println("Directory.ReadData()")
	err = tools.JsonReadData(d.Path + "/" + folder+"/"+filename, data)
	if err != nil {
		return
	}
	return
}


