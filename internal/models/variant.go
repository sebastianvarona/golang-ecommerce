package models

import (
	"time"
)

type Variant struct {
	ID              int       `json:"id"`
	ProductID       int       `json:"product_id"`
	AdditionalPrice float64   `json:"additional_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type VariantTranslation struct {
	ID         int    `json:"id"`
	VariantID  int    `json:"variant_id"`
	LanguageID int    `json:"language_id"`
	Name       string `json:"name"`
}
