package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pong!"))
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/login", HandleLoginPage)
			r.Post("/token", HandleTokenCreation)
			r.Post("/authorize", HandleUserAuthorization)
		})
	})

	log.Println("Serving on :3000 ...")
	http.ListenAndServe(":3000", r)
}
