package controllers

import (
	"fmt"
	"forum/database"
	"forum/models"
	"forum/templ"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func CreatePost(w http.ResponseWriter, r *http.Request, data models.PageData) {
	data.PageTitle = "Create Post"

	if r.Method == http.MethodPost {
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)

		firstFile, fHeader, _ := r.FormFile("image")
		if fHeader != nil {
			defer firstFile.Close()
		}

		categories := r.Form["categories"]
		title := r.FormValue("title")
		content := r.FormValue("content")

		if isEmpty(title) {
			data.Data = "Title must not be empty"
			InternalError(w, r, templ.ExecTemplate(w, "createpost.html", data))
			return
		}
		if isEmpty(content) {
			data.Data = "Content must not be empty"
			InternalError(w, r, templ.ExecTemplate(w, "createpost.html", data))
			return
		}
		if fHeader != nil {
			if fHeader.Size > 20000000 {
				data.Data = "File too large, limit size 20MB"
				InternalError(w, r, templ.ExecTemplate(w, "createpost.html", data))
				return
			}
			if !regex.MatchString(fHeader.Filename) {
				data.Data = "Invalid type, please upload jpg, jpeg, png, gif, svg"
				InternalError(w, r, templ.ExecTemplate(w, "createpost.html", data))
				return
			}
		}
		timeNow := time.Now()
		loc, _ := time.LoadLocation("Asia/Almaty")
		timeNow = timeNow.In(loc)

		newPost := models.Post{
			UserID:     data.User.UserID,
			Username:   data.User.Username,
			Title:      title,
			Content:    strings.ReplaceAll(content, "\n", "<br>"),
			Categories: categories,
			ImageExist: fHeader != nil,
			DateTime:   timeNow,
			TimeString: timeNow.Format("2006-01-02 15:04"),
		}

		err := database.CreatePost(&newPost)
		if InternalError(w, r, err) {
			return
		}

		if fHeader != nil {
			mainFile, _ := os.Create(fmt.Sprintf("./static/images/%v", newPost.PostID))
			defer mainFile.Close()
			_, err = io.Copy(mainFile, firstFile)
			if InternalError(w, r, err) {
				return
			}
		}
	} else if r.Method == http.MethodGet {
		InternalError(w, r, templ.ExecTemplate(w, "createpost.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
func isEmpty(text string) bool {
	for _, r := range text {
		if !(r <= ' ') {
			return false
		}
	}
	return true
}
