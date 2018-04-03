package main

import (
	"net/http"
	"fmt"
	"github.com/vegh1010/moviems/tools"
)

type Handler struct {
	list []RouteDescription
}

func (self *Handler) Init(List []Route) {
	for i := 0; i < len(List); i++ {
		self.list = append(self.list, List[i].Output())
	}
}

func (self *Handler) Help(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler.Help()")
	output := tools.NewJsonOutput(w)

	output.Success = true
	output.Data = self.list
	output.Print()
}
