package main

import (
	"log"
	"net/http"
)

// TokenContext stores the context to a request regarding a token.
type TokenContext struct {
	Username string
	Password string

	RedirectURI string
}

func (token TokenContext) CreateToken() (string, error) {
	// TODO: Check if username and password are good.

	return GenerateRandomString(32)

	// TODO: Store the created token.
}

// HandleTokenCreation handles a POST request to /api/v1/user/token
func HandleTokenCreation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := TokenContext{
		Username:    r.FormValue("username"),
		Password:    r.FormValue("password"),
		RedirectURI: r.FormValue("redirect_uri"),
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

	w.Header().Add("Set-Cookie", "token="+token)
	http.Redirect(w, r, ctx.RedirectURI, http.StatusSeeOther)
}
