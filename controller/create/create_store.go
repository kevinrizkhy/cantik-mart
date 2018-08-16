package create

import (
	//"fmt"
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	//"html/template"
	//"math"
	"net/http"
	//"strconv"
)

//var t *template.Template

func CreateStore(w http.ResponseWriter, r *http.Request) {
	msg := ""
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			address := r.FormValue("address")
			phone := r.FormValue("phone")
			if id != "" && name != "" && address != "" && phone != "" {
				insert_store_status := database.InsertStore(id, name, address, phone)
				if insert_store_status {
					w.WriteHeader(200)
				} else {
					w.WriteHeader(403)
				}
			} else {
				w.WriteHeader(403)
			}
			w.Write([]byte(msg))
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
