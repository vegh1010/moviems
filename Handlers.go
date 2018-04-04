package main

import (
	"github.com/vegh1010/moviems/Param"
	"github.com/vegh1010/moviems/Handlers/CinemaHandler"
	"github.com/vegh1010/moviems/Handlers/FilmHandler"
)

//initialize all handlers
//return Routes defined
func InitHandlers() (Routes) {
	var list Routes
	var Param moviems_param.Param
	var CinemaHandler moviems_cinema_handler.CinemaHandler
	var FilmHandler moviems_film_handler.FilmHandler

	//initialize param
	err := Param.Init()
	check(err)

	//pass param to handlers and initialize
	CinemaHandler.Param = Param
	CinemaHandler.Init()
	FilmHandler.Param = Param
	FilmHandler.Init()

	//add routes information, pattern, description and reference handler's function
	list.Add(
		"CinemaHandler.List()",
		"GET",
		"/v1/cinemaworld",
		"List movies from cinema world",
		CinemaHandler.List,
	)
	list.Add(
		"CinemaHandler.Get()",
		"GET",
		"/v1/cinemaworld/{id}",
		"Get a specific movie from cinema world",
		CinemaHandler.Get,
	)
	list.Add(
		"FilmHandler.List()",
		"GET",
		"/v1/filmworld",
		"List movies from film world",
		FilmHandler.List,
	)
	list.Add(
		"FilmHandler.Get()",
		"GET",
		"/v1/filmworld/{id}",
		"Get a specific movie from film world",
		FilmHandler.Get,
	)
	return list
}

//check if err do panic on error
func check(err error) {
	if err != nil {
		panic(err)
	}
}
