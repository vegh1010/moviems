// package moviems_lib is a moviems library which contains structs/functions related to moviems.
package moviems_lib

// MovieDetail struct used to get detailed information for a movie from api get specific movie by id
type MovieDetail struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	Rated     string `json:"Rated"`
	Released  string `json:"Released"`
	Runtime   string `json:"Runtime"`
	Genre     string `json:"Genre"`
	Director  string `json:"Director"`
	Writer    string `json:"Writer"`
	Actors    string `json:"Actors"`
	Plot      string `json:"Plot"`
	Language  string `json:"Language"`
	Country   string `json:"Country"`
	Awards    string `json:"Awards"`
	Poster    string `json:"Poster"`
	Metascore string `json:"Metascore"`
	Rating    string `json:"Rating"`
	Votes     string `json:"Votes"`
	ID        string `json:"ID"`
	Type      string `json:"Type"`
	Price     string `json:"Price"`
}
