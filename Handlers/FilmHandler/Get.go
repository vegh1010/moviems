// package moviems_film_handler is a handler package for filmword where each function represents to each api pattern
// link to them and handles filmword requests
package moviems_film_handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/vegh1010/moviems/lib"
)

// FilmHandler.List() will get a specific movie from feature Database via filmworld by id as well as handle needed error messages
func (h *FilmHandler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FilmHandler-Get()")

	//get new instance of json output
	output := NewJsonOutput(w)
	vars := mux.Vars(r)

	//get id from path pattern
	id := vars["id"]

	//validate token
	err := moviems_lib.ValidateToken(r, h.AccessToken)
	if output.PrintErrorIf(err, nil, 401) {
		//error handling
		return
	}

	//get movie
	data, err := h.Database.Get(h.FilmWorld, id)
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

