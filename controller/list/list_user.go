package list

import (
	sessions "github.com/pardev/cantik-mart/model/session"
	store "github.com/pardev/cantik-mart/model/store"
	user "github.com/pardev/cantik-mart/model/user"
	"html/template"
	"net/http"
)

func ListUser(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	session_token, id := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/layout.html",
			"view/partial/table/list_user.html",
			"view/partial/base/header/header.html",
			"view/partial/base/navbar/navbar.html",
		)
		data := user.GetUserDetail(id)
		data["list_user"] = template.JS(user.GetUserList())
		data["list_user_role"] = user.GetUserRole()
		data["list_store"] = store.GetStore()
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
