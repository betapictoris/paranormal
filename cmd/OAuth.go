package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type AuthorizationContext struct {
	ResponseType string
	Client       OAuthClient
	RedirectURI  string
	Scope        []string
	State        string

	Code string
}

func HandleUserAuthorization(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx := AuthorizationContext{
		ResponseType: r.FormValue("response_type"),
		RedirectURI:  r.FormValue("redirect_uri"),
		Scope:        strings.Split(r.FormValue("scope"), ","),
		State:        r.FormValue("state"),
	}

	// Try to find the client
	client := clientStore.GetClientByID(r.FormValue("client_id"))
	if client == nil {
		log.Println("Failed to find OAuth 2.0 client:", r.FormValue("client_id"))
		http.Error(w, "OAuth Client Not Found", http.StatusNotFound)
		return
	}

	ctx.Client = *client

	// Create a auth code
	var err error
	ctx.Code, err = GenerateRandomString(16)
	if err != nil {
		log.Println("Failed to generate random string:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	t, err := template.ParseFiles("templates/user/authorize.html")
	if err != nil {
		log.Println("Failed to find template:", "templates/user/authorize.html")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, ctx)
	if err != nil {
		log.Println("Failed to execute template:", err)
	}
}
