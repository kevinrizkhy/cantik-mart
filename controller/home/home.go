package home

import (
	"html/template"
	"net/http"
)

var t *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles(
		"view/layout.html",
	)
	data := make(map[string]interface{})
	t.ExecuteTemplate(w, "layout", data)
}
