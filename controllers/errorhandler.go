package controllers

import (
<<<<<<< HEAD
	"forum/sessions"
=======
>>>>>>> 6d07c3e234a0fecc926f96433a292eddf6b1fe8f
	"net/http"
)

func InternalError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "500 Internal Error")
		return true
	}
	return false
}
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
<<<<<<< HEAD
	user, _ := sessions.GetUser(w, r)
=======
	user
>>>>>>> 6d07c3e234a0fecc926f96433a292eddf6b1fe8f
}
