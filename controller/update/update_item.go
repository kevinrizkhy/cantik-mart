package update

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

var t *template.Template

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		r.ParseForm()
		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			category := r.FormValue("category")
			buy := r.FormValue("buy")
			sell := r.FormValue("sell")
			margin := r.FormValue("margin")
			description := r.FormValue("description")
			if id != "" && name != "" && category != "" && buy != "" && sell != "" && margin != "" {
				buy_int, _ := strconv.Atoi(buy)
				sell_int, _ := strconv.Atoi(sell)
				margin_float_r, _ := strconv.ParseFloat(margin, 64)
				margin_float := math.Ceil(margin_float_r*100) / 100
				buy = fmt.Sprint(buy_int)
				sell = fmt.Sprint(sell_int)
				margin = fmt.Sprint(margin_float)
				update_item_status := database.UpdateItem(id, name, category, buy, sell, margin, description)
				if update_item_status {
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
