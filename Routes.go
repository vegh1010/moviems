package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

//Route struct contains information and referenced handler on routes
type Route struct {
	Name                string
	Method              string
	Pattern             string
	Description         string
	AccessTokenRequired bool
	HandlerFunc         http.HandlerFunc
}

//Route.Output() return RouteDescription
func (r *Route) Output() (data RouteDescription) {
	data.Name = r.Name
	data.Method = r.Method
	data.Pattern = r.Pattern
	data.Description = r.Description
	data.AccessTokenRequired = r.AccessTokenRequired
	return
}

//Routes struct contains a list of Route
type Routes struct {
	list []Route
}

//Route.Add() add new routes information and referenced handler into list
func (r *Routes) Add(Name, Method, Pattern, Description string, HandlerFunc http.HandlerFunc) {
	r.list = append(r.list, Route{
		Name,
		Method,
		Pattern,
		Description,
		true,
		HandlerFunc,
	})
}

// NewRouter create mux.Router based on list of routes
// return mux.Router
func NewRouter(data Routes) *mux.Router {
	//create help handler
	var helpHandler Handler
	//initialize help handler
	helpHandler.Init(data.list)
	helper := Route{
		"Help",
		"GET",
		"/help",
		"List api description on this microservice",
		false,
		helpHandler.Help,
	}
	//added helper handler first to route
	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods(helper.Method).
		Path(helper.Pattern).
		Name(helper.Name).
		Handler(helper.HandlerFunc)

	//added the rest of the handlers
	for _, route := range data.list {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
