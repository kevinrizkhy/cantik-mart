package list

import (
	item "github.com/pardev/cantik-mart/model/item"
	sessions "github.com/pardev/cantik-mart/model/session"
	user "github.com/pardev/cantik-mart/model/user"
	"html/template"
	"net/http"
)

var t *template.Template

func ListItems(w http.ResponseWriter, r *http.Request) {
	session_token, id := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/layout.html",
			"view/partial/table/list_items.html",
			"view/partial/base/header/header.html",
			"view/partial/base/navbar/navbar.html",
		)
		data := user.GetUserDetail(id)
		data["list_items"] = template.JS(item.GetItemList())
		data["list_items_category"] = item.GetItemCategoryListArray()
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
