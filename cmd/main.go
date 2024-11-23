package main

import (
	"html/template"
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
			r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
				t, err := template.ParseFiles("templates/user/login.html")
				if err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}

				err = t.Execute(w, nil)
				if err != nil {
					log.Println("Failed to fully execute template:", err)
				}
			})

			r.Post("/token", func(w http.ResponseWriter, r *http.Request) {
				r.ParseForm()
				ctx := TokenContext{
					Username: r.FormValue("username"),
					Password: r.FormValue("password"),
				}

				token, err := ctx.CreateToken()
				if err != nil {
					log.Println("Failed to generate a user token:", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}

				if token == "" {
					log.Println("Token generation returned an empty token.")
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}

				w.Write([]byte(token))
			})
		})
	})

	log.Println("Serving on :3000 ...")
	http.ListenAndServe(":3000", r)
}
