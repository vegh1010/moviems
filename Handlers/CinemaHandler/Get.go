// package moviems_cinema_handler is a handler package for cinemaworld where each function represents to each api pattern
// link to them and handles cineworld requests
package moviems_cinema_handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/vegh1010/moviems/lib"
)

// CinemaHandler.List() will get a specific movie from feature Database via cinemaworld by id as well as handle needed error messages
func (h *CinemaHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CinemaHandler-Get()")

	//get new instance of json output
	output := NewJsonOutput(w)
	vars := mux.Vars(r)

	//get id from path pattern
	id := vars["id"]

	//validate token
	err := moviems_lib.ValidateToken(r, h.AccessToken)
	if output.PrintErrorIf(err, nil, 403) {
		//error handling
		return
	}

	//get movie
	data, err := h.Database.Get(h.CinemaWorld, id)
	if output.PrintErrorIf(err, nil, 400) {
		//error handling
		return
	}
	if output.PrintIf(data.ID == "", "Movie Not Found", 404) {
		//error handling
		return
	}

	//output result
	output.Data = data
	output.Success = true
	output.Print()
}

