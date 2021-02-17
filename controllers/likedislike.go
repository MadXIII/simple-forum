package controllers

import (
	"forum/database"
	"forum/models"
	"net/http"
	"strconv"
)

func LikePost(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodPost {
		postid, _ := strconv.Atoi(r.FormValue("postid"))
		liked := r.FormValue("submit")
		link := r.FormValue("link")
		var err error

		if (liked != "like" || liked != "dislike") || link == "" {
			ErrorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
			return
		}

		if liked == "like" {
			err = database.LikePost(data.User.UserID, postid)
		} else {
			err = database.DislikePost(data.User.UserID, postid)
		}
		if InternalError(w, r, err) {
			return
		}
		http.Redirect(w, r, link, http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

func LikeComment(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodPost {
		postid, _ := strconv.Atoi(r.FormValue("postid"))
		liked := r.FormValue("submit")
		link := r.FormValue("link")
		var err error

		if (liked != "like" || liked != "dislike") || link == "" {
			ErrorHandler(w, r, http.StatusBadRequest, "400 Bad Request")
			return
		}

		if liked == "like" {
			err = database.LikeComment(data.User.UserID, postid)
		} else {
			err = database.DislikeComment(data.User.UserID, postid)
		}
		if InternalError(w, r, err) {
			return
		}
		http.Redirect(w, r, link, http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
