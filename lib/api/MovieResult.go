// package moviems_api is a moviems library that contains structs.
// The structs will be used to get information from external api's based on code exercise's api defined
package moviems_api

import (
	"github.com/vegh1010/moviems/lib"
	"github.com/vegh1010/moviems/tools"
	"encoding/json"
	"errors"
)

// MovieResult struct is used to get a specific movie
type MovieResult struct {
	moviems_lib.MovieDetail
}

// MovieResult.Get() will accept url, movie provider, id and access token which will be used to retrieve a specific
// movie from external api. It will unmarshal the response into itself.
// It returns error if issue occurred.
func (m *MovieResult) Get(url, from, id, accessToken string) (error) {
	//format url
	url = url + "/api/" + from + "/movie/" + id

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
