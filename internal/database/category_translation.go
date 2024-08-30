package database

import (
	"database/sql"
	"ecommerce_template/internal/models"
)

func (s *service) CreateCategoryTranslation(ct *models.CategoryTranslation) error {
	query := `
        INSERT INTO category_translations (category_id, language_id, name, description)
        VALUES ($1, $2, $3, $4) RETURNING id`
	return s.db.QueryRow(query, ct.CategoryID, ct.LanguageID, ct.Name, ct.Description).Scan(&ct.ID)
}

func (s *service) GetCategoryTranslationByID(id int) (*models.CategoryTranslation, error) {
	ct := &models.CategoryTranslation{}
	query := `
        SELECT id, category_id, language_id, name, description
        FROM category_translations WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&ct.ID, &ct.CategoryID, &ct.LanguageID, &ct.Name, &ct.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return ct, err
}

func (s *service) GetAllCategoryTranslations() ([]*models.CategoryTranslation, error) {
	query := `SELECT id, category_id, language_id, name, description FROM category_translations`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var translations []*models.CategoryTranslation
	for rows.Next() {
		ct := &models.CategoryTranslation{}
		err := rows.Scan(&ct.ID, &ct.CategoryID, &ct.LanguageID, &ct.Name, &ct.Description)
		if err != nil {
			return nil, err
		}
		translations = append(translations, ct)
	}

	return translations, nil
}

func (s *service) UpdateCategoryTranslation(ct *models.CategoryTranslation) error {
	query := `
        UPDATE category_translations SET name = $1, description = $2
        WHERE id = $3`
	_, err := s.db.Exec(query, ct.Name, ct.Description, ct.ID)
	return err
}

func (s *service) DeleteCategoryTranslation(id int) error {
	query := `DELETE FROM category_translations WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
