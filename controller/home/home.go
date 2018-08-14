package home

import (
	sessions "github.com/pardev/cantik-mart/model/session"
	user "github.com/pardev/cantik-mart/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func Home(w http.ResponseWriter, r *http.Request) {
	session_token, id := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/layout.html",
			"view/partial/base/alert/alert.html",
			"view/partial/base/card/card.html",
			"view/partial/base/chart/chart.html",
			"view/partial/base/header/header.html",
			"view/partial/base/navbar/navbar.html",
		)
		data := user.GetUserDetail(id)
		data["name1"] = "qqweqweqwe"
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
