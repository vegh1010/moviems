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
	CheckEnvironment()

	port := os.Getenv("PORT")

	log.SetOutput(os.Stdout)
	now := time.Now()
	fmt.Println("Server starting at ", now.Format("2006-01-02 15:04:05"), " on port: "+port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, NewRouter(InitHandlers()))
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
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
