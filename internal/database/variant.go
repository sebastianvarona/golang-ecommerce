package database

import (
	"database/sql"
	"ecommerce_template/internal/models"
	"time"
)

func (s *service) CreateVariant(variant *models.Variant) error {
	query := `
        INSERT INTO variants (product_id, additional_price, created_at, updated_at)
        VALUES ($1, $2, $3, $4) RETURNING id`
	return s.db.QueryRow(query, variant.ProductID, variant.AdditionalPrice, time.Now(), time.Now()).Scan(&variant.ID)
}

func (s *service) GetVariantByID(id int) (*models.Variant, error) {
	variant := &models.Variant{}
	query := `
        SELECT id, product_id, additional_price, created_at, updated_at
        FROM variants WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&variant.ID, &variant.ProductID, &variant.AdditionalPrice, &variant.CreatedAt, &variant.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return variant, err
}

func (s *service) GetAllVariants() ([]*models.Variant, error) {
	query := `SELECT id, product_id, additional_price, created_at, updated_at FROM variants`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []*models.Variant
	for rows.Next() {
		variant := &models.Variant{}
		err := rows.Scan(&variant.ID, &variant.ProductID, &variant.AdditionalPrice, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			return nil, err
		}
		variants = append(variants, variant)
	}

	return variants, nil
}

func (s *service) UpdateVariant(variant *models.Variant) error {
	query := `
        UPDATE variants SET product_id = $1, additional_price = $2, updated_at = $3
        WHERE id = $4`
	_, err := s.db.Exec(query, variant.ProductID, variant.AdditionalPrice, time.Now(), variant.ID)
	return err
}

func (s *service) DeleteVariant(id int) error {
	query := `DELETE FROM variants WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
