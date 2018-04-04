// package moviems_param provides a struct that contains variables and functions needed to start moviems.
// It will initialize variables from static values and environment variables.
// It will create directories required for caching purpose.
package moviems_param

import (
	"github.com/vegh1010/moviems/tools"
	"fmt"
)

// Param.CreateDirectories() will create cinemaworld and filmworld folder in cache folder
// it returns an error if issue occurred
func (p *Param) CreateDirectories() (err error) {
	fmt.Println("Param.CreateDirectories()")

	//list of path directories
	list := []string{
		p.Path + "/" + p.CinemaWorld,
		p.Path + "/" + p.FilmWorld,
	}

	for _, data := range list {
		//create directory
		err = tools.CreateDirectory(data)
		if err != nil {
			return
		}
	}
	return
}

