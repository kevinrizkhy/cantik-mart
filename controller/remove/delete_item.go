package remove

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"html/template"
	"net/http"
)

var t *template.Template

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			category := r.FormValue("category")
			if id != "" && name != "" && category != "" {
				delete_item_status := database.DeleteItem(id, name, category)
				if delete_item_status {
					w.WriteHeader(200)
				} else {
					w.WriteHeader(403)
				}
			} else {
				w.WriteHeader(403)
			}
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
