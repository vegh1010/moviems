package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name                string
	Method              string
	Pattern             string
	Description         string
	AccessTokenRequired bool
	HandlerFunc         http.HandlerFunc
}

func (self *Route) Output() (data RouteDescription) {
	data.Name = self.Name
	data.Method = self.Method
	data.Pattern = self.Pattern
	data.Description = self.Description
	data.AccessTokenRequired = self.AccessTokenRequired
	return
}

type Routes struct {
	list []Route
}

func (self *Routes) Add(Name, Method, Pattern, Description string, HandlerFunc http.HandlerFunc) {
	self.list = append(self.list, Route{
		Name,
		Method,
		Pattern,
		Description,
		true,
		HandlerFunc,
	})
}

func NewRouter(data Routes) *mux.Router {
	var helpHandler Handler
	helpHandler.Init(data.list)
	helper := Route{
		"Help",
		"GET",
		"/help",
		"List api description on this microservice",
		false,
		helpHandler.Help,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods(helper.Method).
		Path(helper.Pattern).
		Name(helper.Name).
		Handler(helper.HandlerFunc)
	for _, route := range data.list {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
