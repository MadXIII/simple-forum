package database

import (
	"forum/models"
)

func CreateComment(newComment *models.Comment) error {
	comment, err := db.Prepare(`
		INSERT INTO comments
		(postid, username, text, imageexist, date_time, timestring)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}

	res, err := comment.Exec(
		newComment.PostID,
		newComment.Username,
		newComment.Text,
		newComment.ImageExist,
		newComment.DateTime,
		newComment.TimeString,
	)
	if err != nil {
		return err
	}

	commentid, _ := res.LastInsertId()
	newComment.CommentID = int(commentid)

	return err
}

func getPostCommentCount(post *models.Post) error {
	if err := db.QueryRow(`
		SELECT *
		FROM comments
		WHERE post = ?
	`, post.PostID).Scan(&post.CommentCount); err != nil {
		return err
	}
	return err
}

func GetCommentsByPostID(postid int, uid int) ([]models.Comment, error) {
	var comments []models.Comment

	row, err := db.Query(`
		SELECT *
		FROM comments
		WHERE postid = ?
	`, postid)
	defer row.Close()
	if err != nil {
		return comments, err
	}

	for row.Next() {
		var comment models.Comment
		row.Scan(&comment.CommentID, &comment.PostID, &comment.Username, &comment.Text, &comment.ImageExist, &comment.DateTime, &comment.TimeString)
		err = getCommentLikesDislikes(&comment)
		if err != nil {
			return comments, err
		}
		if uid != 0 {
			err = commentLikedByUser(&comment, uid)
			if err != nil {
				return comments, err
			}
			err = commentDislikedByUser(&comment, uid)
			if err != nil {
				return comments, err
			}
		}
		comments = append(comments, comment)
	}
	return comments, err
}
