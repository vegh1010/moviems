// package moviems_lib is a moviems library which contains structs/functions related to moviems.
package moviems_lib

// Movie struct used to get information from list api
type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ID     string `json:"ID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
