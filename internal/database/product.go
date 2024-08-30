package database

import (
	"database/sql"
	"ecommerce_template/internal/models"
	"time"
)

func (s *service) CreateProduct(product *models.Product) error {
	query := `
        INSERT INTO products (category_id, price, stock, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return s.db.QueryRow(query, product.CategoryID, product.Price, product.Stock, time.Now(), time.Now()).Scan(&product.ID)
}

func (s *service) GetProductByID(id int) (*models.Product, error) {
	product := &models.Product{}
	query := `
        SELECT id, category_id, price, stock, created_at, updated_at
        FROM products WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&product.ID, &product.CategoryID, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return product, err
}

func (s *service) GetAllProducts() ([]*models.Product, error) {
	query := `SELECT id, category_id, price, stock, created_at, updated_at FROM products`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.CategoryID, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *service) UpdateProduct(product *models.Product) error {
	query := `
        UPDATE products SET category_id = $1, price = $2, stock = $3, updated_at = $4
        WHERE id = $5`
	_, err := s.db.Exec(query, product.CategoryID, product.Price, product.Stock, time.Now(), product.ID)
	return err
}

func (s *service) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}
