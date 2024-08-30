package database

import (
	"ecommerce_template/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func chackPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *service) UserInsert(user *models.User) (int, error) {
	fail := func(err error) (int, error) { return 0, fmt.Errorf("UserInsert: %v", err) }
	hash, err := hashPassword(user.Password)
	if err != nil {
		return fail(err)
	}
	query := `INSERT INTO users (password_hash, first_name, last_name, email, address, phone_number, locale) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
	var id int
	err = s.db.QueryRow(query, hash, user.FirstName, user.LastName, user.Email, user.Address, user.PhoneNumber, user.Locale).Scan(&id)
	if err != nil {
		return fail(err)
	}
	return id, nil
}
