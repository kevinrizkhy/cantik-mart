package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

func SetSession(w http.ResponseWriter, r *http.Request, session_value string) {
	var store = sessions.NewCookieStore([]byte("%SESSION%2104%"))
	sessionToken, _ := store.Get(r, "session-token")
	sessionToken.Values["token"] = session_value
	sessionToken.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24,
		HttpOnly: true,
	}
	expiration := time.Now().Add(7 * 24 * time.Hour)
	cookiename := http.Cookie{Name: "kXpIV89wIQUxA86p2uWIAuOfflmGNZyYP", Value: session_value, Expires: expiration}
	http.SetCookie(w, &cookiename)
	sessions.Save(r, w)
}
