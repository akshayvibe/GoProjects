package routes

import (
	"gowithpg/internal/handler"
	"net/http"
)

// server mux is a multiplexer which redirects the incoming request to the particular function

func RegisterRoutes(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/stocks", h.CreateStock)
	mux.HandleFunc("GET /api/stocks", h.GetAllStock)
	mux.HandleFunc("GET /api/stocks/{id}", h.GetStock)
	mux.HandleFunc("PUT /api/stocks/{id}", h.UpdateStock)
	mux.HandleFunc("DELETE /api/stocks/{id}", h.DeleteStock)

	return mux
}