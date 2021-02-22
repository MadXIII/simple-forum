package controllers

import (
	"fmt"
	"forum/database"
	"forum/models"
	"net/http"
	"strconv"
	"time"
)

func CreateComment(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodPost {

		text := r.FormValue("text")
		postid, _ := strconv.Atoi(r.FormValue("postid"))

		if isEmpty(text) {
			ErrorHandler(w, r, http.StatusUnprocessableEntity, "Text must not be empty")
			return
		}

		timeNow := time.Now()
		loc, _ := time.LoadLocation("Asia/Almaty")
		timeNow = timeNow.In(loc)

		newComment := models.Comment{
			PostID:     postid,
			Username:   data.User.Username,
			Text:       text,
			DateTime:   timeNow,
			TimeString: timeNow.Format("2006-01-02 15:04"),
		}

		err := database.CreateComment(&newComment)
		if InternalError(w, r, err) {
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/posts/id//%d#%d", newComment.PostID, newComment.CommentID), http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
