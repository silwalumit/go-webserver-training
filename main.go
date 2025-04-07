package main

import (
	"log"
	"net/http"
	"server/middleware"
	"server/router"
)

func main() {
	router := router.NewRouter()

	// add middlewares
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Finding by id: " + id))
	})
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Starting server on port :8000")
	server.ListenAndServe()
}
