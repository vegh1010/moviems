// package moviems_param provides a struct that contains variables and functions needed to start moviems.
// It will initialize variables from static values and environment variables.
// It will create directories required for caching purpose.
package moviems_param

import (
	"os"
)

// Param contains variables required for moviems
type Param struct {
	URL         string //url for external api
	AccessToken string //access token to do request from external api
	Path        string //directory of starting path
	CinemaWorld string //cinema world folder name
	FilmWorld   string //film world folder name
}

// Param.Init() will initialize value and call CreateDirectories()
func (p *Param) Init() (err error) {
	//get environment variables
	p.URL = os.Getenv("URL")
	p.AccessToken = os.Getenv("ACCESS_TOKEN")

	//set static values
	p.Path = "cache"
	p.CinemaWorld = "cinemaworld"
	p.FilmWorld = "filmworld"

	//create directories
	err = p.CreateDirectories()
	if err != nil {
		return
	}
	return
}
