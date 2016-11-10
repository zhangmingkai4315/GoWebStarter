package controllers

import (
	"net/http"
	"fmt"
)

func NotExistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotExist 404")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
}
