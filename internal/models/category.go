package models

import (
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	ParentID  *int      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryTranslation struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	LanguageID  int    `json:"language_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
