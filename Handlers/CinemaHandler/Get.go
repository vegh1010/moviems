package moviems_cinema_handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/vegh1010/moviems/lib"
)

func (h *CinemaHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CinemaHandler-Get()")
	output := NewJsonOutput(w)
	vars := mux.Vars(r)
	id := vars["id"]

	err := moviems_lib.ValidateToken(r, h.AccessToken)
	if output.PrintErrorIf(err, nil) {
		return
	}

	data, err := h.Database.Get(h.CinemaWorld, id)
	if output.PrintErrorIf(err, nil) {
		return
	}
	if output.PrintIf(data.ID == "", "Movie Not Found") {
		return
	}

	output.Data = data
	output.Success = true
	output.Print()
}

