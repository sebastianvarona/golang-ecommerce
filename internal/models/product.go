package models

import (
	"time"
)

type Product struct {
	ID         int       `json:"id"`
	CategoryID *int      `json:"category_id"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductTranslation struct {
	ID          int    `json:"id"`
	ProductID   int    `json:"product_id"`
	LanguageID  int    `json:"language_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
