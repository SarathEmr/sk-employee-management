package app

import (
	"github.com/go-chi/chi/v5"
)

func (server Server) InitRouter() chi.Router {
	r := chi.NewRouter()

	r.Route("/sk-em/employee", func(r chi.Router) {
		r.Get("/", server.handleListEmployees)
	})
	return r
}
