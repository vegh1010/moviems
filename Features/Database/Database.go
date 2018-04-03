package moviems_database

import (
	"github.com/vegh1010/moviems/Param"
	"github.com/vegh1010/moviems/Features/Directory"
)

type Database struct {
	moviems_param.Param
	Directory moviems_directory.Directory
}

func (d *Database) Init() {
	d.Directory.Param = d.Param
	d.Directory.Init()
}
