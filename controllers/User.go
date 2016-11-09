package controllers

import (
	"fmt"
	"net/http"

	"views"
)

// get /login handler
func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", "base", struct{}{})
}

// post /login handler
func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post login")
}

// get /SignUp handler
func GetSignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SignUp page")
}

// post /SignUp handler
func PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post sing up")
}
