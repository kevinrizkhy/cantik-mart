package update

import (
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			number := r.FormValue("number")
			order_by := r.FormValue("order_by")
			order_to := r.FormValue("order_to")
			total := r.FormValue("total")
			note := r.FormValue("note")
			status := r.FormValue("status")
			if id != "" && number != "" && total != "" && order_by != "" && order_to != "" && status != "" && note != "" {
				update_store_status := database.UpdateOrder(id, number, order_by, order_to, total, note, status)
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
