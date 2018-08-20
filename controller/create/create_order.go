package create

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	msg := ""
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			number := r.FormValue("number")
			total := r.FormValue("total")
			order_by := r.FormValue("order_by")
			sent_to := r.FormValue("sent_to")
			note := r.FormValue("note")
			if number != "" && total != "" && order_by != "" && sent_to != "" && note != "" {
				insert_order_status := database.InsertOrder(number, total, order_by, sent_to, note)
				if insert_order_status {
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
