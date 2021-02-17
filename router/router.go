package router

import (
	"forum/controllers"
	"net/http"
)

func ServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./static/images"))))
	mux.HandleFunc("/", controllers.MainPaige)
	mux.HandleFunc("/register", PageDataMiddleWare(false, controllers.Register))
	mux.HandleFunc("/login", PageDataMiddleWare(false, controllers.Login))
	mux.HandleFunc("/logut", controllers.Logout)
	mux.HandleFunc("/posts", PageDataMiddleWare(false, controllers.GetPosts))
	mux.HandleFunc("/posts/", PageDataMiddleWare(false, controllers.GetPostsByCategory))
	mux.HandleFunc("/posts/id/", PageDataMiddleWare(false, controllers.GetPostByID))
	mux.HandleFunc("/createposts", PageDataMiddleWare(true, controllers.CreatePost))
	mux.HandleFunc("/myposts", PageDataMiddleWare(true, controllers.GetMyPosts))
	mux.HandleFunc("/mylikes", PageDataMiddleWare(true, controllers.GetMyLikedPosts))
	mux.HandleFunc("/comment", PageDataMiddleWare(true, controllers.CreateComment))
	mux.HandleFunc("/likepost", PageDataMiddleWare(true, controllers.LikePost))
	mux.HandleFunc("/likecomment", PageDataMiddleWare(true, controllers.LikeComment))
	return mux
}
