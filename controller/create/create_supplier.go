package create

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func CreateSupplier(w http.ResponseWriter, r *http.Request) {
	msg := ""
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			name := r.FormValue("name")
			address := r.FormValue("address")
			phone := r.FormValue("phone")
			email := r.FormValue("email")
			if name != "" && address != "" && phone != "" && email != "" {
				insert_supplier_status := database.InsertSupplier(name, address, phone, email)
				if insert_supplier_status {
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
