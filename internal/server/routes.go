package server

import (
	"ecommerce_template/cmd/web"
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)

	app := web.NewApplication(s.db)

	r.Group(func(r chi.Router) { // pages routes, optional auth
		r.Get("/", app.IndexPageHandler)
		r.Get("/register", app.RegisterPageHandler)
		r.Post("/register", app.RegisterHandler)
	})

	r.Group(func(r chi.Router) { // auth routes
		r.Get("/admin", app.DashboardPageHandler)
		r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
		r.Post("/hello", web.HelloWebHandler)
	})

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
