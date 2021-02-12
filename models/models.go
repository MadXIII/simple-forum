package models

import (
	"html/template"
	"time"
)

type User struct {
	UserID   int
	Salt     string
	Hash     []byte
	Username string
	Email    string
}

type Post struct {
	PostID     int
	UserID     int
	Username   string
	Title      string
	Content    template.HTML
	Categories []string
	// CommentCount int
	DateTime   time.Time
	TimeString string
	Like       int
	Dislike    int
	Liked      bool
	Disliked   bool
}
type Session struct {
	SessionID   string
	UserID      int
	TimeCreated time.Time
}
type PageData struct {
	PageTitle  string
	Categories []string
	User       User
	Data       interface{}
}
type Comment struct {
	CommentID  int
	PostID     int
	Username   string
	Text       template.HTML
	DateTime   time.Time
	TimeString string
	Like       int
	Dislike    int
	Liked      bool
	Disliked   bool
}
