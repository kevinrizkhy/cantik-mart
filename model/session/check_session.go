package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	database "github.com/pardev/cantik-mart/model/db"
	"net/http"
)

func CheckSession(r *http.Request) (bool, string) {
	var store = sessions.NewCookieStore([]byte("%SESSION%2104%"))
	sessionToken, _ := store.Get(r, "session-token")
	session_val := sessionToken.Values["token"]
	rows, err := database.ExecuteQuery("SELECT * FROM t_user WHERE session_token = $1 AND status=$2;", fmt.Sprint(session_val), true)
	if len(rows) > 0 && err == nil {
		return true, fmt.Sprint(rows[0]["id"])
	} else {
		if err != nil {
			fmt.Println("ERR : CheckSession - ", err)
		}
		return false, ""
	}

}
