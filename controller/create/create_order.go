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

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	msg := ""
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			name := r.FormValue("name")
			supplier := r.FormValue("supplier")
			total := r.FormValue("total")
			note := r.FormValue("note")
			if name != "" && supplier != "" && total != "" && note != "" {
				/*buy_int, _ := strconv.Atoi(buy)
				sell_int, _ := strconv.Atoi(sell)
				margin_float_r, _ := strconv.ParseFloat(margin, 64)
				margin_float := math.Ceil(margin_float_r*100) / 100
				buy = fmt.Sprint(buy_int)
				sell = fmt.Sprint(sell_int)
				margin = fmt.Sprint(margin_float)*/
				insert_order_status := database.InsertOrder(name, supplier, total, note)
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
