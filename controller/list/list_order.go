package list

import (
	item "github.com/pardev/cantik-mart/model/item"
	order "github.com/pardev/cantik-mart/model/order"
	sessions "github.com/pardev/cantik-mart/model/session"
	//store "github.com/pardev/cantik-mart/model/store"
	supplier "github.com/pardev/cantik-mart/model/supplier"
	user "github.com/pardev/cantik-mart/model/user"
	"html/template"
	"net/http"
)

func ListOrder(w http.ResponseWriter, r *http.Request) {
	session_token, id := sessions.CheckSession(r)
	if session_token {
		t, _ = template.ParseFiles(
			"view/layout.html",
			"view/partial/table/list_order.html",
			"view/partial/base/header/header.html",
			"view/partial/base/navbar/navbar.html",
		)
		data := user.GetUserDetail(id)
		data["list_order"] = template.JS(order.GetOrder())
		data["list_items"] = template.JS(item.GetItemList())
		data["list_supplier"] = supplier.GetSupplier()
		//data["list_items_category"] = item.GetItemCategoryListArray()

		/*data["list_user"] = template.JS(user.GetUserList())
		data["list_store"] = store.GetStore()
		*/

		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
