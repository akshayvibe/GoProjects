package handler

import (
	"encoding/json"
	"gowithpg/internal/db"
	"gowithpg/internal/model"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	DB     db.StockStore
	Logger *slog.Logger
}

// Constructor
func New(db db.StockStore, logger *slog.Logger) *Handler {
	return &Handler{
		DB:     db,
		Logger: logger,
	}
}

// HEALTH CHECK
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {

	h.Logger.Info("health check endpoint hit")

	w.Header().Set("Content-Type", "application/json")

	err := h.DB.Ping()

	if err != nil {

		h.Logger.Error("database ping failed",
			slog.String("error", err.Error()),
		)

		w.WriteHeader(http.StatusServiceUnavailable)

		json.NewEncoder(w).Encode(map[string]string{
			"status": "database down",
		})

		return
	}

	h.Logger.Info("health check successful")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// CREATE STOCK
func (h *Handler) CreateStock(w http.ResponseWriter, r *http.Request) {

	h.Logger.Info("create stock request received")

	var stock model.Stock

	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {

		h.Logger.Error("failed to decode request body",
			slog.String("error", err.Error()),
		)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateStock(&stock); err != nil {

		h.Logger.Error("failed to create stock",
			slog.String("error", err.Error()),
			slog.String("symbol", stock.Symbol),
		)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info("stock created successfully",
		slog.Uint64("id", uint64(stock.ID)),
		slog.String("symbol", stock.Symbol),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// GET ONE STOCK
func (h *Handler) GetStock(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {

		h.Logger.Warn("invalid stock id",
			slog.String("id", mux.Vars(r)["id"]),
		)

		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	h.Logger.Info("fetching stock",
		slog.Int("id", id),
	)

	stock, err := h.DB.GetStock(id)

	if err != nil {

		h.Logger.Error("stock fetch failed",
			slog.Int("id", id),
			slog.String("error", err.Error()),
		)

		http.Error(w, "stock not found", http.StatusNotFound)
		return
	}

	h.Logger.Info("stock fetched successfully",
		slog.Int("id", id),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// GET ALL STOCKS
func (h *Handler) GetAllStock(w http.ResponseWriter, r *http.Request) {

	h.Logger.Info("fetching all stocks")

	stocks, err := h.DB.GetAllStocks()

	if err != nil {

		h.Logger.Error("failed to fetch stocks",
			slog.String("error", err.Error()),
		)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info("stocks fetched successfully",
		slog.Int("count", len(stocks)),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

// UPDATE STOCK
func (h *Handler) UpdateStock(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {

		h.Logger.Warn("invalid stock id for update",
			slog.String("id", mux.Vars(r)["id"]),
		)

		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var stock model.Stock

	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {

		h.Logger.Error("failed to decode update request",
			slog.String("error", err.Error()),
		)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.UpdateStock(id, &stock); err != nil {

		h.Logger.Error("failed to update stock",
			slog.Int("id", id),
			slog.String("error", err.Error()),
		)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stock.ID = uint(id)

	h.Logger.Info("stock updated successfully",
		slog.Int("id", id),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

// DELETE STOCK
func (h *Handler) DeleteStock(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {

		h.Logger.Warn("invalid stock id for delete",
			slog.String("id", mux.Vars(r)["id"]),
		)

		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.DB.DeleteStock(id); err != nil {

		h.Logger.Error("failed to delete stock",
			slog.Int("id", id),
			slog.String("error", err.Error()),
		)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info("stock deleted successfully",
		slog.Int("id", id),
	)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(id)
}