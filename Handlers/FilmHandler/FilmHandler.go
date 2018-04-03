package moviems_file_handler

import (
	"github.com/vegh1010/moviems/Param"
	"github.com/vegh1010/moviems/Features/Database"
	"github.com/vegh1010/moviems/tools"
)

var NewJsonOutput = tools.NewJsonOutput

type FilmHandler struct {
	moviems_param.Param
	Database moviems_database.Database
}

func (h *FilmHandler) Init() {
	h.Database.Param = h.Param
	h.Database.Init()
}
