package web

import (
	"ecommerce_template/cmd/web/pages"
	"ecommerce_template/internal/models"
	"net/http"
	"strings"
)

func (app *application) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		serverError(w, err)
		return
	}
	user := models.User{
		FirstName:   r.PostForm.Get("first_name"),
		LastName:    r.PostForm.Get("last_name"),
		Email:       strings.TrimSpace(r.PostForm.Get("email")),
		Address:     r.PostForm.Get("address"),
		PhoneNumber: r.PostForm.Get("phone_number"),
		Locale:      "en",
	}
	_, err = app.db.UserInsert(&user)
	if err != nil {
		serverError(w, err)
		return
	}
	app.IndexPageHandler(w, r)
}

func (app *application) DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	pages.DashboardPage().Render(r.Context(), w)
}
