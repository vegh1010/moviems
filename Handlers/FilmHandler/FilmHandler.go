// package moviems_film_handler is a handler package for filmword where each function represents to each api pattern
// link to them and handles filmword requests
package moviems_film_handler

import (
	"github.com/vegh1010/moviems/Param"
	"github.com/vegh1010/moviems/Features/Database"
	"github.com/vegh1010/moviems/tools"
)

//reference NewJsonOutput for standard json output for handler
var NewJsonOutput = tools.NewJsonOutput

// FilmHandler struct is created for easier access to functions available and contain
// variables required for moviems from moviems_param.Param.
// Feature Database is part of it as a variable which allows it to have access to
// Database's variables and functions.
type FilmHandler struct {
	moviems_param.Param
	Database moviems_database.Database
}

// FilmHandler.Init() will be used to pass param and initialize any feature packages that contains on itself.
func (h *FilmHandler) Init() {
	//initialize Database
	h.Database.Param = h.Param
	h.Database.Init()
}
