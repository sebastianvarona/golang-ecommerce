package database

import (
	"database/sql"
	"ecommerce_template/internal/models"
)

func (s *service) CreateProductTranslation(pt *models.ProductTranslation) error {
	query := `
        INSERT INTO product_translations (product_id, language_id, name, description)
        VALUES ($1, $2, $3, $4) RETURNING id`
	return s.db.QueryRow(query, pt.ProductID, pt.LanguageID, pt.Name, pt.Description).Scan(&pt.ID)
}

func (s *service) GetProductTranslationByID(id int) (*models.ProductTranslation, error) {
	pt := &models.ProductTranslation{}
	query := `
        SELECT id, product_id, language_id, name, description
        FROM product_translations WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&pt.ID, &pt.ProductID, &pt.LanguageID, &pt.Name, &pt.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return pt, err
}

func (s *service) GetAllProductTranslations() ([]*models.ProductTranslation, error) {
	query := `SELECT id, product_id, language_id, name, description FROM product_translations`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var translations []*models.ProductTranslation
	for rows.Next() {
		pt := &models.ProductTranslation{}
		err := rows.Scan(&pt.ID, &pt.ProductID, &pt.LanguageID, &pt.Name, &pt.Description)
		if err != nil {
			return nil, err
		}
		translations = append(translations, pt)
	}

	return translations, nil
}

func (s *service) UpdateProductTranslation(pt *models.ProductTranslation) error {
	query := `
        UPDATE product_translations SET name = $1, description = $2
        WHERE id = $3`
	_, err := s.db.Exec(query, pt.Name, pt.Description, pt.ID)
	return err
}

func (s *service) DeleteProductTranslation(id int) error {
	query := `DELETE FROM product_translations WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
