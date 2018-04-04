// package moviems_cinema_handler is a handler package for cinemaworld where each function represents to each api pattern
// link to them and handles cineworld requests
package moviems_cinema_handler

import (
	"net/http"
	"fmt"
	"github.com/vegh1010/moviems/lib"
)

// CinemaHandler.List() will get a list of movies from feature Database via cinemaworld as well as handle needed error messages
func (h *CinemaHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CinemaHandler-List()")

	//get new instance of json output
	output := NewJsonOutput(w)

	//validate token
	err := moviems_lib.ValidateToken(r, h.AccessToken)
	if output.PrintErrorIf(err, nil, 403) {
		//error handling
		return
	}

	//get movies
	data, err := h.Database.List(h.CinemaWorld)
	if output.PrintErrorIf(err, nil, 400) {
		//error handling
		return
	}

	//output result
	output.Data = data
	output.Success = true
	output.Print()
}

