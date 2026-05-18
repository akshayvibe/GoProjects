package model

import "time"

type Stock struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Symbol    string    `gorm:"size:20;unique;not null" json:"symbol"`
	Price     float64   `gorm:"not null" json:"price"`
	Quantity  int       `gorm:"default:0" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}