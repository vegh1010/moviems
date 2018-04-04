// package moviems_database is a feature package database which handles getting data from either from external apis
// or from cache files.
// It will handle writing data to files if data retrieve successfully.
// If it fails to get data from external apis, it will check if there's a cache data version of the request.
package moviems_database

import (
	"github.com/vegh1010/moviems/Param"
	"github.com/vegh1010/moviems/Features/Directory"
)

// Database struct is created for easier access to functions available and contain
// variables required for moviems from moviems_param.Param.
// Feature Directory is part of Database as a variable which allows it to have access to
// Directory's variables and functions.
type Database struct {
	moviems_param.Param
	Directory moviems_directory.Directory
}

// Database.Init() will be used to pass param and initialize any feature packages that contains on itself.
func (d *Database) Init() {
	//initialize Directory
	d.Directory.Param = d.Param
	d.Directory.Init()
}
