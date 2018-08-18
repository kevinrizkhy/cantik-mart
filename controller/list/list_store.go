package list

import (
	sessions "github.com/pardev/cantik-mart/model/session"
	store "github.com/pardev/cantik-mart/model/store"
	user "github.com/pardev/cantik-mart/model/user"
	"html/template"
	"net/http"
)

func ListStore(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	session_token, id := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/layout.html",
			"view/partial/table/list_store.html",
			"view/partial/base/header/header.html",
			"view/partial/base/navbar/navbar.html",
		)
		data := user.GetUserDetail(id)
		data["list_store"] = store.GetStore()
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
