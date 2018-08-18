package update

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			email := r.FormValue("email")
			address := r.FormValue("address")
			phone := r.FormValue("phone")
			role := r.FormValue("role")
			store := r.FormValue("store")
			status := r.FormValue("status")
			if id != "" && name != "" && email != "" && address != "" && phone != "" && role != "" && store != "" && status != "" {
				update_user_status := database.UpdateUserDetail(id, name, email, address, phone, role, store, status)
				if update_user_status {
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
