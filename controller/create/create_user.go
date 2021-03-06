package create

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	msg := ""
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			name := r.FormValue("name")
			email := r.FormValue("email")
			address := r.FormValue("address")
			phone := r.FormValue("phone")
			role := r.FormValue("role")
			store := r.FormValue("store")
			password := r.FormValue("password")
			if name != "" && email != "" && address != "" && phone != "" && password != "" && role != "" && store != "" {
				insert_user_status := database.InsertUser(name, email, address, phone, role, store, password)
				if insert_user_status {
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
