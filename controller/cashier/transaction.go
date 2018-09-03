package cashier

import (
	sessions "github.com/pardev/cantik-mart/model/session"
	"html/template"
	"net/http"
)

var t *template.Template

func Transaction(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/cashier/transaction.html",
		)
		// data := user.GetUserDetail(id)
		t.ExecuteTemplate(w, "layout", nil)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
