package sign

import (
	sessions "github.com/pardev/cantik-mart/model/session"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if session_token {
		sessions.ClearSession(w, r)
	}
	http.Redirect(w, r, "/login", 302)
}
