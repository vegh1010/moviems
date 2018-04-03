package moviems_database

import (
	"github.com/vegh1010/moviems/lib"
	"github.com/vegh1010/moviems/lib/api"
	"fmt"
)

func (d *Database) List(from string) (list []moviems_lib.Movie, err error) {
	fmt.Println("Database.List()")
	var api moviems_api.MoviesResult
	err = api.List(d.URL, from, d.AccessToken)
	if err != nil {
		if d.Directory.IsFileExist(from, from + ".json") {
			err = d.Directory.ReadData(from, from + ".json", &list)
			if err != nil {
				return
			}
		}
		return
	}
	list = api.Movies

	if d.Directory.IsFileExist(from, from + ".json") {
		err = d.Directory.DeleteFile(from, from + ".json")
		if err != nil {
			return
		}
	}
	err = d.Directory.CreateFile(from, from + ".json")
	if err != nil {
		return
	}
	err = d.Directory.WriteData(from, from + ".json", list)
	if err != nil {
		return
	}
	return
}
