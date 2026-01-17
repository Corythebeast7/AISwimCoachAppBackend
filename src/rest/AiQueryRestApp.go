package rest

import (
	"AISwimCoachAppBackend/src/service"
	"github.com/gorilla/mux"
	"net/http"
)

func initAiQueryRoutes(r *mux.Router) {

	r.HandleFunc("/api/ai/query").Methods("POST")

}

func aiQuery(w http.ResponseWriter, r *http.Request) {
	service.AiQueryHandler(w, r)
}
