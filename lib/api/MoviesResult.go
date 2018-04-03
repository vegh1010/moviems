package moviems_api

import (
	"github.com/vegh1010/moviems/lib"
	"github.com/vegh1010/moviems/tools"
	"encoding/json"
	"errors"
)

type MoviesResult struct {
	Movies []moviems_lib.Movie `json:"Movies"`
}

func (m *MoviesResult) List(url, from, accessToken string) (error) {
	url = url + "/api/" + from + "/movies"
	result, err := tools.GetRequest(url, accessToken)
	if err != nil {
		return err
	}
	if result.StatusCode != 200 {
		return errors.New(string(result.ResponseBody))
	}

	err = json.Unmarshal(result.ResponseBody, m)
	if err != nil {
		return err
	}
	return nil
}
