package router

import (
	//"fmt"
	"github.com/gorilla/mux"
	api_controller "github.com/pardev/cantik-mart/controller/api"
	cashier_controller "github.com/pardev/cantik-mart/controller/cashier"
	create_controller "github.com/pardev/cantik-mart/controller/create"
	home_controller "github.com/pardev/cantik-mart/controller/home"
	list_controller "github.com/pardev/cantik-mart/controller/list"
	delete_controller "github.com/pardev/cantik-mart/controller/remove"
	sign_controller "github.com/pardev/cantik-mart/controller/sign"
	update_controller "github.com/pardev/cantik-mart/controller/update"
	"net/http"
)

func Routes(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/sign", sign)

	r.HandleFunc("/cashier", cashier)

	r.HandleFunc("/list/{type}", list)

	r.HandleFunc("/api/{type}", api)

	r.HandleFunc("/create/{type}", create)
	r.HandleFunc("/edit/{type}", edit)
	r.HandleFunc("/delete/{type}", deleteFunc)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/assets/", r)
	http.ListenAndServe(addr, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	sign_controller.Login(w, r)
}

func logout(w http.ResponseWriter, r *http.Request) {
	sign_controller.Logout(w, r)
}

func home(w http.ResponseWriter, r *http.Request) {
	home_controller.Home(w, r)
}

func sign(w http.ResponseWriter, r *http.Request) {
	sign_controller.SignIn(w, r)
}

func cashier(w http.ResponseWriter, r *http.Request) {
	cashier_controller.Transaction(w, r)
}

func api(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			api_controller.GetItem(w, r)
		} else if types == "item-list" {
			api_controller.GetItemList(w, r)
		} else if types == "member" {
			//api_controller.GetMember(w, r)
		}
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			list_controller.ListItems(w, r)
		} else if types == "user" {
			list_controller.ListUser(w, r)
		} else if types == "store" {
			list_controller.ListStore(w, r)
		} else if types == "member" {
			list_controller.ListMember(w, r)
		} else if types == "order" {
			list_controller.ListOrder(w, r)
			/*} else if types == "orderItem" {
			list_controller.ListOrderItem(w, r)*/
		} else if types == "supplier" {
			list_controller.ListSupplier(w, r)
		}
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			create_controller.CreateItem(w, r)
		} else if types == "user" {
			create_controller.CreateUser(w, r)
		} else if types == "store" {
			create_controller.CreateStore(w, r)
		} else if types == "order" {
			create_controller.CreateOrder(w, r)
		} else if types == "supplier" {
			create_controller.CreateSupplier(w, r)
		}

	}
}

func edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			update_controller.UpdateItem(w, r)
		} else if types == "user" {
			update_controller.UpdateUser(w, r)
		} else if types == "store" {
			update_controller.UpdateStore(w, r)
		} else if types == "supplier" {
			update_controller.UpdateSupplier(w, r)
		}
	}
}

func deleteFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			delete_controller.DeleteItem(w, r)
		} else if types == "store" {
		}
	}
}
