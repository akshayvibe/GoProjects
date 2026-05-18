package routes

import (
	"gowithpg/internal/handler"
	// "net/http"

	"github.com/gorilla/mux"
)

// server mux is a multiplexer which redirects the incoming request to the particular function

func RegisterRoutes(h *handler.Handler) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/health", h.HealthCheck).Methods("GET")

	router.HandleFunc("/api/stocks", h.CreateStock).Methods("POST")
	router.HandleFunc("/api/stocks", h.GetAllStock).Methods("GET")

	router.HandleFunc("/api/stocks/{id}", h.GetStock).Methods("GET")
	router.HandleFunc("/api/stocks/{id}", h.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stocks/{id}", h.DeleteStock).Methods("DELETE")

	return router
}