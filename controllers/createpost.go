package controllers

import (
	"fmt"
	"forum/database"
	"forum/models"
	"forum/templ"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func CreatePost(w http.ResponseWriter, r *http.Request, data models.PageData) {
	data.PageTitle = "Create Post"

	if r.Method == http.MethodPost {
		firstFile, fHeader, _ := r.FormFile("image")
		if firstFile != nil {
			defer firstFile.Close()
		}

		categories := r.Form["categories"]
		title := r.FormValue("title")
		content := r.FormValue("content")
		categoryExist := make(map[string]bool)

		for _, category := range data.Categories {
			categoryExist[category] = true
		}

		for _, category := range categories {
			if !categoryExist[category] {
				data.Data = "Invalid category " + category
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
				return
			}
		}
		if isEmpty(title) {
			data.Data = "Title must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
			return
		}

		if !isValidLenOfTitle(title) {
			data.Data = "Title must be between 2-60 characters"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
			return
		}
		if isEmpty(content) {
			data.Data = "Content must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
			return
		}
		if firstFile != nil {
			imgData, err := ioutil.ReadAll(firstFile)
			if err != nil {
				data.Data = "Cannot upload image"
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
				return
			}
			if fHeader.Size > 20971520 {
				data.Data = "File too large, limit size 20MB"
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
				return
			}
			firstFile.Seek(0, 0)
			correctType := []string{"image/png", "image/gif", "image/jpeg"}
			contType := http.DetectContentType(imgData)

			for _, imgType := range correctType {
				if imgType != contType {
					data.Data = "Invalid type, please upload jpeg, png, gif"
					InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
					w.WriteHeader(http.StatusUnprocessableEntity)
					return
				}
				break
			}

		}
		timeNow := time.Now()
		loc, _ := time.LoadLocation("Asia/Almaty")
		timeNow = timeNow.In(loc)

		newPost := models.Post{
			UserID:     data.User.UserID,
			Username:   data.User.Username,
			Title:      title,
			Content:    content,
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
		http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		InternalError(w, r, templ.ExecuteTemplate(w, "createpost.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
func isEmpty(text string) bool {
	for _, r := range text {
		if !(r <= 32) {
			return false
		}
	}
	return true
}

func isValidLenOfTitle(title string) bool {
	if len(title) < 2 || len(title) > 60 {
		return false
	}
	return true
}
