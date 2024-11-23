package main

import (
	"html/template"
	"log"
	"net/http"
)

// LoginContext is the context given to the login template.
type LoginContext struct {
	RedirectURI string
}

// HandleLoginPage handles a request for the user to login.
func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := LoginContext{
		RedirectURI: r.Form.Get("redirect_uri"),
	}

	t, err := template.ParseFiles("templates/user/login.html")
	if err != nil {
		log.Println("Failed to find template:", "templates/user/login.html")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, ctx)
	if err != nil {
		log.Println("Failed to fully execute template:", err)
	}
}
