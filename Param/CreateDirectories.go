package moviems_param

import "github.com/vegh1010/moviems/tools"

//create cinemaworld and filmworld folder
func (p *Param) CreateDirectories() (err error) {
	list := []string{
		p.Path + "/" + p.CinemaWorld,
		p.Path + "/" + p.FilmWorld,
	}

	for _, data := range list {
		err = tools.CreateDirectory(data)
		if err != nil {
			return
		}
	}
	return
}

