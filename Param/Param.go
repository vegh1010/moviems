package moviems_param

import (
	"os"
)

//varaibles in the struct are common variables used by all packages in moviems
type Param struct {
	URL         string
	AccessToken string
	Path        string
	CinemaWorld string
	FilmWorld   string
}

//initialize value
func (p *Param) Init() (err error) {
	p.URL = os.Getenv("URL")
	p.AccessToken = os.Getenv("ACCESS_TOKEN")

	p.Path = "./"
	p.CinemaWorld = "cinemaworld"
	p.FilmWorld = "filmworld"

	err = p.CreateDirectories()
	if err != nil {
		return
	}
	return
}
