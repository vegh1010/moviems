package moviems_file_handler

import (
	"net/http"
	"fmt"
	"github.com/vegh1010/moviems/lib"
)

func (h *FilmHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FilmHandler-List()")
	output := NewJsonOutput(w)

	err := moviems_lib.ValidateToken(r, h.AccessToken)
	if output.PrintErrorIf(err, nil) {
		return
	}

	data, err := h.Database.List(h.FilmWorld)
	if output.PrintErrorIf(err, nil) {
		return
	}

	output.Data = data
	output.Success = true
	output.Print()
}

