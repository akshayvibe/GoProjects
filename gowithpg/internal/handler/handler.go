package handler

import (
	"encoding/json"
	// storage "gowithpg/internal/db/postgres"
	// storage "gowithpg/internal/db/postgres"
	"gowithpg/internal/db"
	"gowithpg/internal/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	DB db.StockStore
}

// CREATE STOCK
func (h *Handler) CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock model.Stock

	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateStock(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// GET ONE STOCK
func (h *Handler) GetStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	stock, err := h.DB.GetStock(id)
	if err != nil {
		http.Error(w, "stock not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// GET ALL STOCKS
func (h *Handler) GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := h.DB.GetAllStocks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

// UPDATE STOCK
func (h *Handler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var stock model.Stock

	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.UpdateStock(id, &stock); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stock.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// DELETE STOCK
func (h *Handler) DeleteStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.DB.DeleteStock(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}