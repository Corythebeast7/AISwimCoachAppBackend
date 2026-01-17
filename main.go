package main

import (
	"AISwimCoachAppBackend/src/rest"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Setup API rest
	rest.SetupRoutes(r)

	// Configure server
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server
	fmt.Println("Server starting on port 8080...")
	log.Fatal(srv.ListenAndServe())
}
