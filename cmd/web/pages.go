package web

import (
	"ecommerce_template/cmd/web/pages"
	"net/http"
)

func (app *application) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	pages.IndexPage().Render(r.Context(), w)
}

func (app *application) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	pages.RegisterPage().Render(r.Context(), w)
}
