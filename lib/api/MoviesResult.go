// package moviems_api is a moviems library that contains structs.
// The structs will be used to get information from external api's based on code exercise's api defined
package moviems_api

import (
	"github.com/vegh1010/moviems/lib"
	"github.com/vegh1010/moviems/tools"
	"encoding/json"
	"errors"
)

// MoviesResult struct is used to get a list of movies
type MoviesResult struct {
	Movies []moviems_lib.Movie `json:"Movies"`
}

// MovieResult.List() will accept url, movie provider and access token which will be used to retrieve a list of movies
// from external api. It will unmarshal the response into itself.
// It returns error if issue occurred.
func (m *MoviesResult) List(url, from, accessToken string) (error) {
	//format url
	url = url + "/api/" + from + "/movies"

	//call get request
	result, err := tools.GetRequest(url, accessToken)
	if err != nil {
		return err
	}

	//check if status code is 200
	if result.StatusCode != 200 {
		return errors.New(string(result.ResponseBody))
	}

	//unmarshal response to itself
	err = json.Unmarshal(result.ResponseBody, m)
	if err != nil {
		return err
	}
	return nil
}
