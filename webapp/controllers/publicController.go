package controllers

import (
	"net/http"
	"webapp/views"
)

func GetIndexPageHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "index", "base", struct{}{})
	return
}