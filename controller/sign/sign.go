package sign

import (
	sessions "github.com/pardev/cantik-mart/model/session"
	signin "github.com/pardev/cantik-mart/model/sign"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if !session_token {
		r.ParseForm()
		if r.Method == "POST" {
			if len(r.FormValue("email")) == 0 {
				http.Redirect(w, r, "/sign", 302)
				return
			}
			email := r.FormValue("email")
			if len(r.FormValue("password")) == 0 {
				http.Redirect(w, r, "/sign", 302)
				return
			}
			password := r.FormValue("password")
			stat := signin.Login(w, r, email, password)
			if stat {
				http.Redirect(w, r, "/", 302)
			} else {
				http.Redirect(w, r, "/login", 302)
			}
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}

}
