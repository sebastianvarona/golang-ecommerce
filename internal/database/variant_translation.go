package database

import (
	"database/sql"
	"ecommerce_template/internal/models"
)

func (s *service) CreateVariantTranslation(vt *models.VariantTranslation) error {
	query := `
        INSERT INTO variant_translations (variant_id, language_id, name)
        VALUES ($1, $2, $3) RETURNING id`
	return s.db.QueryRow(query, vt.VariantID, vt.LanguageID, vt.Name).Scan(&vt.ID)
}

func (s *service) GetVariantTranslationByID(id int) (*models.VariantTranslation, error) {
	vt := &models.VariantTranslation{}
	query := `
        SELECT id, variant_id, language_id, name
        FROM variant_translations WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&vt.ID, &vt.VariantID, &vt.LanguageID, &vt.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return vt, err
}

func (s *service) GetAllVariantTranslations() ([]*models.VariantTranslation, error) {
	query := `SELECT id, variant_id, language_id, name FROM variant_translations`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var translations []*models.VariantTranslation
	for rows.Next() {
		vt := &models.VariantTranslation{}
		err := rows.Scan(&vt.ID, &vt.VariantID, &vt.LanguageID, &vt.Name)
		if err != nil {
			return nil, err
		}
		translations = append(translations, vt)
	}

	return translations, nil
}

func (s *service) UpdateVariantTranslation(vt *models.VariantTranslation) error {
	query := `
        UPDATE variant_translations SET name = $1
        WHERE id = $2`
	_, err := s.db.Exec(query, vt.Name, vt.ID)
	return err
}

func (s *service) DeleteVariantTranslation(id int) error {
	query := `DELETE FROM variant_translations WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
