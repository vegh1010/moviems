// package moviems_film_handler is a handler package for filmword where each function represents to each api pattern
// link to them and handles filmword requests
package moviems_film_handler

import (
	"net/http"
	"fmt"
	"github.com/vegh1010/moviems/lib"
)

// FilmHandler.List() will get a list of movies from feature Database via filmworld as well as handle needed error messages
func (h *FilmHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FilmHandler-List()")

	//get new instance of json output
	output := NewJsonOutput(w)

	//validate token
	err := moviems_lib.ValidateToken(r, h.AccessToken)
	if output.PrintErrorIf(err, nil, 403) {
		//error handling
		return
	}

	//get movies
	data, err := h.Database.List(h.FilmWorld)
	if output.PrintErrorIf(err, nil, 400) {
		//error handling
		return
	}

	//output result
	output.Data = data
	output.Success = true
	output.Print()
}

