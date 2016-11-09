package views

import (
	"html/template"
	"net/http"
)

var templates map[string](*template.Template)

func init() {
	if templates == nil {
		templates = make(map[string](*template.Template))
	}
	templates["index"] = template.Must(template.ParseFiles("views/index.html", "views/base.html"))
	templates["login"] = template.Must(template.ParseFiles("views/login.html", "views/base.html"))
	templates["signup"] = template.Must(template.ParseFiles("views/signup.html", "views/base.html"))
}

func RenderTemplate(w http.ResponseWriter, name string, temp string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if ok != false {
		http.Error(w, "The url is not exist", http.StatusNotFound)
	}
	err := tmpl.ExecuteTemplate(w, temp, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
