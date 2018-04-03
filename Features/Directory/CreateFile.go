package moviems_directory

import (
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

//create a file in the specified folder
func (d *Directory) CreateFile(folder, filename string) (err error) {
	fmt.Println("Directory.CreateFile()")
	_, err = tools.CreateFile(d.Path+"/"+folder, filename)
	if err != nil {
		return
	}
	return
}
