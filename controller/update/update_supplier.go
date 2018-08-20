package update

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			address := r.FormValue("address")
			phone := r.FormValue("phone")
			email := r.FormValue("email")
			status := r.FormValue("status")
			if id != "" && name != "" && address != "" && phone != "" && email != "" && status != "" {
				update_store_status := database.UpdateSupplier(id, name, address, phone, email, status)
				if update_store_status {
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
