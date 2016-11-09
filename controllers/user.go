package controllers

import (
	"fmt"
	"net/http"

	"views"
	"models"
	"time"
)

// get /login handler
func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "login", "base", struct{}{})
}

// post /login handler
func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "post login")
	var user models.User
	r.ParseForm()
	user.Username=r.PostFormValue("username")
	user.Password=r.PostFormValue("password")
	user.CreatedAt=time.Now()
	fmt.Fprintf(w,"%v",user)
}


// get /SignUp handler
func GetSignUpHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "signup", "base", struct{}{})
}

// post /SignUp handler
func PostSignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post sing up")
}
