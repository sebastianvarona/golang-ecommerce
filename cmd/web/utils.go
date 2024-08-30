package web

import (
	"ecommerce_template/internal/database"
	"fmt"
	"net/http"
)

type application struct {
	db database.Service
}

func NewApplication(db database.Service) *application {
	return &application{db: db}
}

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("server error: %v", err), 500)
}
