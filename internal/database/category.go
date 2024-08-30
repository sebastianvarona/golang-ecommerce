package database

import (
	"database/sql"
	"ecommerce_template/internal/models"
	"time"
)

func (s *service) CreateCategory(category *models.Category) error {
	query := `
        INSERT INTO categories (parent_id, created_at, updated_at)
        VALUES ($1, $2, $3) RETURNING id`
	return s.db.QueryRow(query, category.ParentID, time.Now(), time.Now()).Scan(&category.ID)
}

func (s *service) GetCategoryByID(id int) (*models.Category, error) {
	category := &models.Category{}
	query := `
        SELECT id, parent_id, created_at, updated_at
        FROM categories WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&category.ID, &category.ParentID, &category.CreatedAt, &category.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return category, err
}

func (s *service) GetAllCategories() ([]*models.Category, error) {
	query := `SELECT id, parent_id, created_at, updated_at FROM categories`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		category := &models.Category{}
		err := rows.Scan(&category.ID, &category.ParentID, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *service) UpdateCategory(category *models.Category) error {
	query := `
        UPDATE categories SET parent_id = $1, updated_at = $2
        WHERE id = $3`
	_, err := s.db.Exec(query, category.ParentID, time.Now(), category.ID)
	return err
}

func (s *service) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
