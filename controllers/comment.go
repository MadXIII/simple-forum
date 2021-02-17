package controllers

import (
	"fmt"
	"forum/database"
	"forum/models"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CreateComment(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodPost {
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)
		firstFile, fileHeader, _ := r.FormFile("image")
		if fileHeader != nil {
			defer firstFile.Close()
		}

		text := r.FormValue("text")
		postid, _ := strconv.Atoi(r.FormValue("postid"))

		if isEmpty(text) {
			ErrorHandler(w, r, http.StatusUnprocessableEntity, "Text must not be empty")
			return
		}
		if fileHeader != nil {
			if fileHeader.Size > 20000000 {
				ErrorHandler(w, r, http.StatusUnprocessableEntity, "File too large, limit size 20MB")
				return
			}
			if !regex.MatchString(fileHeader.Filename) {
				ErrorHandler(w, r, http.StatusUnprocessableEntity, "Invalid type, please upload jpg, jpeg, png, gif, svg")
				return
			}
		}
		timeNow := time.Now()
		loc, _ := time.LoadLocation("Asia/Almaty")
		timeNow = timeNow.In(loc)

		newComment := models.Comment{
			PostID:     postid,
			Username:   data.User.Username,
			Text:       strings.ReplaceAll(text, "\n", "<br>"),
			ImageExist: fileHeader != nil,
			DateTime:   timeNow,
			TimeString: timeNow.Format("2006-01-02 15:04"),
		}

		err := database.CreateComment(&newComment)
		if InternalError(w, r, err) {
			return
		}

		if fileHeader != nil {
			file, _ := os.Create(fmt.Sprintf("./static/images/c%v", newComment.CommentID))
			defer file.Close()
			_, err = io.Copy(file, firstFile)
			if InternalError(w, r, err) {
				return
			}
		}
		http.Redirect(w, r, fmt.Sprintf("/posts/id//%d#%d", newComment.PostID, newComment.CommentID), http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
