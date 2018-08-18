package list

import (
	"fmt"
	sessions "github.com/pardev/cantik-mart/model/session"
	supplier "github.com/pardev/cantik-mart/model/supplier"
	user "github.com/pardev/cantik-mart/model/user"
	"html/template"
	"net/http"
)

func ListSupplier(w http.ResponseWriter, r *http.Request) {
	session_token, id := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/layout.html",
			"view/partial/table/list_items.html",
			"view/partial/base/header/header.html",
			"view/partial/base/navbar/navbar.html",
		)
		data := user.GetUserDetail(id)
		data["list_supplier"] = template.JS(supplier.GetSupplierList())
		fmt.Println(data)
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
