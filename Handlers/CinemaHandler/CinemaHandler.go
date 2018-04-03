package moviems_cinema_handler

import (
	"github.com/vegh1010/moviems/Param"
	"github.com/vegh1010/moviems/Features/Database"
	"github.com/vegh1010/moviems/tools"
)

var NewJsonOutput = tools.NewJsonOutput

type CinemaHandler struct {
	moviems_param.Param
	Database moviems_database.Database
}

func (h *CinemaHandler) Init() {
	h.Database.Param = h.Param
	h.Database.Init()
}
