package sign

import (
	config "github.com/pardev/cantik-mart/config"
	sessions "github.com/pardev/cantik-mart/model/session"
	"html/template"
	"net/http"
)

var t *template.Template

func Login(w http.ResponseWriter, r *http.Request) {
	session_token, _ := sessions.CheckSession(r)
	if !session_token {
		t, _ = template.ParseFiles(
			"view/sign/sign.html",
		)
		data := make(map[string]interface{})
		data["BaseURL"] = config.Base_URL
		t.ExecuteTemplate(w, "layout", data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
