package db

import "gowithpg/internal/model"
type StockStore interface {
	CreateStock(*model.Stock) error
	GetStock(int) (*model.Stock, error)
	GetAllStocks() ([]model.Stock, error)
	UpdateStock(int, *model.Stock) error
	DeleteStock(int) error
	Ping()error
}