package sign

import (
	config "github.com/wellcode/pardev/proposal/config"
	"html/template"
	"net/http"
)

var t *template.Template

func Sign(w http.ResponseWriter, r *http.Request) {
	t, _ = template.ParseFiles(
		"view/sign/sign.html",
	)
	data := make(map[string]interface{})
	data["BaseURL"] = config.Base_URL
	t.ExecuteTemplate(w, "layout", data)
}
