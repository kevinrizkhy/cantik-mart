package router

import (
	home "github.com/wellcode/pardev/proposal/controller/home"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	home.Home(w, r)
}
