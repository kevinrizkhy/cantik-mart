package sign

import (
	"fmt"
	database "github.com/pardev/cantik-mart/model/db"
	sessions "github.com/pardev/cantik-mart/model/session"
	utility "github.com/pardev/cantik-mart/model/utility"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, email string, password string) bool {
	rows, err := database.ExecuteQuery("SELECT * from t_user where email=$1 and password=$2", email, password)
	if len(rows) > 0 && err == nil {
		sessionToken := utility.GenerateToken(email + password)
		sessions.SetSession(w, r, sessionToken)
		_, err := database.ExecuteQuery("update t_user set session_token=$1 where email=$2", sessionToken, email)
		if err == nil {
			return true
		} else {
			fmt.Println("Login - 1", err)
			return false
		}
	}
	return false
}
