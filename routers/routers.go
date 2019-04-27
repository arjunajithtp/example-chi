package routers

import (
	"net/http"

	"github.com/arjunajithtp/example-chi/handlers"
	"github.com/arjunajithtp/example-chi/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router is for implementing chi
func Router() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/id", func(r chi.Router) {
		r.Use(middlewares.CustomAuth)
		r.Get("/{id}", handlers.GetID)
		r.Post("/{id}", handlers.GetAll)
	})

	http.ListenAndServe(":1234", r)
}
