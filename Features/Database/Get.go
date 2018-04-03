package moviems_database

import (
	"github.com/vegh1010/moviems/lib"
	"fmt"
	"github.com/vegh1010/moviems/lib/api"
)

func (d *Database) Get(from, id string) (data moviems_lib.MovieDetail, err error) {
	fmt.Println("Database.Get()")
	var api moviems_api.MovieResult
	err = api.Get(d.URL, from, id, d.AccessToken)
	if err != nil {
		if d.Directory.IsFileExist(from, id + ".json") {
			err = d.Directory.ReadData(from, id + ".json", &data)
			if err != nil {
				return
			}
		}
		return
	}
	data = api.MovieDetail

	if d.Directory.IsFileExist(from, id + ".json") {
		err = d.Directory.DeleteFile(from, id + ".json")
		if err != nil {
			return
		}
	}
	err = d.Directory.CreateFile(from, id + ".json")
	if err != nil {
		return
	}
	err = d.Directory.WriteData(from, id + ".json", data)
	if err != nil {
		return
	}
	return
}
