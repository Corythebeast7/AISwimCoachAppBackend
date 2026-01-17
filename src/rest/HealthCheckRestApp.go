package rest

import (
	"AISwimCoachAppBackend/src/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func initHealthCheckRoutes(r *mux.Router) {

	r.HandleFunc("/api/healthcheck", HealthCheck).Methods("GET")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	healthCheck := service.HealthCheck()
	if healthCheck.Status == "OK" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonBytes, _ := json.Marshal(healthCheck)
	w.Write(jsonBytes)
}
