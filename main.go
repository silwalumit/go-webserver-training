package main

import (
	"log"
	"net/http"
	"server/handlers"
	"server/middleware"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.HelloWorld)

	server := &http.Server{
		Addr:    ":8000",
		Handler: middleware.LoggingMiddleware(router),
	}

	log.Println("Starting server on port :8000")
	server.ListenAndServe()
}
