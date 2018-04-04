// package moviems_directory is a feature package directory which handles anything related to directory functions
// such as create, write, read and delete file as well as check if file exist.
package moviems_directory

import "github.com/vegh1010/moviems/Param"

// Directory struct is created for easier access to functions available and contain
// variables required for moviems from moviems_param.Param
type Directory struct {
	moviems_param.Param
}

// Directory.Init() will be used to pass param and initialize any feature packages that contains on itself.
func (d *Directory) Init() {

}
