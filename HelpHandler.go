package main

import (
	"net/http"
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

//Handler struct contains a list of apis available for moviems
type Handler struct {
	list []RouteDescription
}

//Handler.Init() initialize list of routes
func (self *Handler) Init(List []Route) {
	for i := 0; i < len(List); i++ {
		self.list = append(self.list, List[i].Output())
	}
}

//Handler.Help() output list of routes available for moviems
func (self *Handler) Help(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler.Help()")

	//get new instance of json output
	output := tools.NewJsonOutput(w)

	//output result
	output.Success = true
	output.Data = self.list
	output.Print()
}

//RouteDescription struct contains information on routes
type RouteDescription struct {
	Name                string `json:"name"`                  //Which handler is reference to this route
	Method              string `json:"method"`                //Request Methods (GET, POST, PATCH, PUT, OPINION)
	Pattern             string `json:"pattern"`               //api pattern defined
	Description         string `json:"description"`           //short description on what this api does
	AccessTokenRequired bool   `json:"access_token_required"` //does this api require access token
}
