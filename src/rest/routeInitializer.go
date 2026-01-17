package rest

import (
	"github.com/gorilla/mux"
)

// SetupRoutes configures all the rest for the API
func SetupRoutes(r *mux.Router) {

	initHealthCheckRoutes(r)
	initAiQueryRoutes(r)

}
