package controllers

import (
	"net/http"
	"webapp/views"
)


// get /admin handler
func AdminIndexHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "admin", "base", struct{}{})
}
