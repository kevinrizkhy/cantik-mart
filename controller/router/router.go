package router

import (
	home "github.com/wellcode/pardev/proposal/controller/home"
	sign "github.com/wellcode/pardev/proposal/controller/sign"
	"net/http"
)

func Sign(w http.ResponseWriter, r *http.Request) {
	sign.Sign(w, r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	home.Home(w, r)
}
