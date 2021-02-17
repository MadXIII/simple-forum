package sessions

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session")
	if err != nil {
		return
	}
	ck.MaxAge = -1
	http.SetCookie(w, ck)
}
