package router

import (
	//"fmt"
	"github.com/gorilla/mux"
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

	r.HandleFunc("/list/{type}", list)

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
		}

	}
}

func edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			update_controller.UpdateItem(w, r)
		}
	}
}

func deleteFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	types := vars["type"]
	if types != "" {
		if types == "item" {
			delete_controller.DeleteItem(w, r)
		}
	}
}
