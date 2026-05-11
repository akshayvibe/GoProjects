package routes

import (
	"gowithpg/internal/handler"
	"net/http"
)

// server mux is a multiplexer which redirects the incoming request to the particular function
func router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /api/stock/{id}",handler.GetStock)
	router.HandleFunc("POST /api/stock/", handler.GetAllStock)
	router.HandleFunc("POST /api/stock/newstock", handler.CreateStock)
	router.HandleFunc("POST /api/stock/{id}", handler.UpdateStock)
	router.HandleFunc("POST /api/stock/deletestock", handler.DeleteStock)
	return router
}