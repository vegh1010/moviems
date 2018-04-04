// package moviems_database is a feature package database which handles getting data from either from external apis
// or from cache files.
// It will handle writing data to files if data retrieve successfully.
// If it fails to get data from external apis, it will check if there's a cache data version of the request.
package moviems_database

import (
	"github.com/vegh1010/moviems/lib"
	"fmt"
	"github.com/vegh1010/moviems/lib/api"
)

// Database.Get() will get a specific movie from external api or from cached data by id
func (d *Database) Get(from, id string) (data moviems_lib.MovieDetail, err error) {
	fmt.Println("Database.Get()")
	var api moviems_api.MovieResult

	//request information from external api
	err = api.Get(d.URL, from, id, d.AccessToken)
	if err != nil {
		//check if file exist
		if d.Directory.IsFileExist(from, id + ".json") {
			//get data from file
			err = d.Directory.ReadData(from, id + ".json", &data)
			if err != nil {
				return
			}
			return
		}
		return
	}
	data = api.MovieDetail

	//write data to file
	err = d.Directory.WriteData(from, id + ".json", data)
	if err != nil {
		return
	}
	return
}
