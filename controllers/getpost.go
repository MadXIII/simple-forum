package controllers

import (
	"forum/database"
	"forum/models"
	"forum/templ"
	"net/http"
	"path"
	"strconv"
)

type postData struct {
	Link  string
	Posts []models.Post
}

func GetPosts(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {
		posts, err := database.GetPosts(data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = "All Posts"
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, templ.ExecTemplate(w, "posts.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

func GetPostByID(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {
		dir, endpoint := path.Split(r.URL.Path)
		postid, _ := strconv.Atoi(endpoint)

		if dir != "/posts/id/" || postid == 0 {
			ErrorHandler(w, r, http.StatusNotFound, "404 Not Found")
			return
		}

		post, err := database.GetPostByID(postid, data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		if post.PostID != postid {
			ErrorHandler(w, r, http.StatusNotFound, "404 Not Found")
			return
		}

		comments, err := database.GetCommentsByPostID(postid, data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = post.Title
		data.Data = struct {
			Post     models.Post
			Comments []models.Comment
		}{
			Post:     post,
			Comments: comments,
		}

		InternalError(w, r, templ.ExecTemplate(w, "post.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

func GetPostsByCategory(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {
		dir, category := path.Split(r.URL.Path)

		if dir != "/posts/" {
			ErrorHandler(w, r, http.StatusNotFound, "404 Not Found")
			return
		}

		for i := range data.Categories {
			if category == data.Categories[i] {
				break
			}
			if i == len(data.Categories)-1 {
				ErrorHandler(w, r, http.StatusNotFound, "404 Not Found")
				return
			}
		}
		posts, err := database.GetPostsByCategory(category, data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = category
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, templ.ExecTemplate(w, "posts.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

func GetMyPosts(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {
		posts, err := database.GetPostsByUserID(data.User.UserID)
		if InternalError(w, r, err) {
			return
		}
		data.PageTitle = "My Posts"
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, templ.ExecTemplate(w, "posts.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}

func GetMyLikedPosts(w http.ResponseWriter, r *http.Request, data models.PageData) {
	if r.Method == http.MethodGet {

		posts, err := database.GetLikedPostsByUserID(data.User.UserID)
		if InternalError(w, r, err) {
			return
		}

		data.PageTitle = "My Likes"
		data.Data = postData{r.URL.Path, posts}
		InternalError(w, r, templ.ExecTemplate(w, "posts.html", data))

	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
