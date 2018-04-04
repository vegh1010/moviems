package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	fmt.Println("MovieMS Loading")

	//check environment
	CheckEnvironment()

	//get port
	port := os.Getenv("PORT")

	//set log output
	log.SetOutput(os.Stdout)

	//set start up time
	now := time.Now()

	//initialize moviems's handlers
	movieHandlers := InitHandlers()

	//get moviems's routes
	movieRoutes := NewRouter(movieHandlers)

	fmt.Println("Server starting at ", now.Format("2006-01-02 15:04:05"), " on port: "+port)

	//set handlers
	loggedRouter := handlers.LoggingHandler(os.Stdout, movieRoutes)

	//define cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"127.0.0.1",
		},
		AllowedMethods:   []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"content-type", "x-access-token"},
	})
	handler := c.Handler(loggedRouter)

	//start listener and server
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
