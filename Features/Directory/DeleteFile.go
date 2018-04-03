package moviems_directory

import (
	"fmt"
	"os"
)

//delete a file in the specified folder
func (d *Directory) DeleteFile(folder, filename string) (err error) {
	fmt.Println("Directory.DeleteFile()")
	err = os.Remove(d.Path + "/" + folder + "/" + filename)
	if err != nil {
		return
	}
	return
}
