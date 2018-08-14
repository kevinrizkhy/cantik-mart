package create

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"html/template"
	"net/http"
)

var t *template.Template

func CreateItem(w http.ResponseWriter, r *http.Request) {
	msg := ""
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("email")
			category := r.FormValue("category")
			buy := r.FormValue("buy")
			sell := r.FormValue("sell")
			margin := r.FormValue("margin")
			description := r.FormValue("description")
			if id != "" && name != "" && category != "" && buy != "" && sell != "" && margin != "" {
				insert_item_status := database.InsertItem(id, name, category, buy, sell, margin, description)
				if insert_item_status {
					msg = "Berhasil menambahkan item."
				} else {
					msg = "Gagal menambahkan item."
				}
			} else {
				w.WriteHeader(403)
				msg = "Gagal menambahkan item."
			}
			w.Write([]byte(msg))
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
