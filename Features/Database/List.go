// package moviems_database is a feature package database which handles getting data from either from external apis
// or from cache files.
// It will handle writing data to files if data retrieve successfully.
// If it fails to get data from external apis, it will check if there's a cache data version of the request.
package moviems_database

import (
	"github.com/vegh1010/moviems/lib"
	"github.com/vegh1010/moviems/lib/api"
	"fmt"
)

// Database.List() will get a list of movies from external api or from cached data
func (d *Database) List(from string) (list []moviems_lib.Movie, err error) {
	fmt.Println("Database.List()")
	var api moviems_api.MoviesResult

	//request information from external api
	err = api.List(d.URL, from, d.AccessToken)
	if err != nil {
		//check if file exist
		if d.Directory.IsFileExist(from, from + ".json") {
			//get list from file
			err = d.Directory.ReadData(from, from + ".json", &list)
			if err != nil {
				return
			}
			return
		}
		return
	}
	list = api.Movies

	//write list to file
	err = d.Directory.WriteData(from, from + ".json", list)
	if err != nil {
		return
	}
	return
}
