package moviems_directory

import (
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

//write data in specified file
func (d *Directory) WriteData(folder, filename string, data interface{}) (err error) {
	fmt.Println("Directory.WriteData()")
	err = tools.WriteData(d.Path + "/" + folder+"/"+filename, data)
	if err != nil {
		return
	}
	return
}


