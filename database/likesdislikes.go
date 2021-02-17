package database

import (
	"forum/models"
)

func LikePost(uid int, postid int) error {
	like, err := db.Prepare(`
		INSERT INTO postlikes
		(userid, postid, liked)
		VALUES (?, ?, 1);
	`)
	if err != nil {
		return err
	}
	liked := 2
	if err = db.QueryRow(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ?)
	`, uid, postid).Scan(&uid, &postid, &liked); err != nil {
		return err
	}
	if liked == 0 {
		like, err = db.Prepare(`
			UPDATE postlikes
			SET liked = 1
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 1 {
		like, err = db.Prepare(`
			DELETE FROM postlikes
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	}
	_, err = like.Exec(
		uid,
		postid,
	)
	if err != nil {
		return err
	}
	return err
}

func DislikePost(uid int, postid int) error {
	dislike, err := db.Prepare(`
		INSERT INTO postlikes
		(userid, postid, liked)
		VALUES (?, ?, 0);
	`)
	if err != nil {
		return err
	}

	liked := 2
	if err = db.QueryRow(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ?)
	`, uid, postid).Scan(&uid, &postid, &liked); err != nil {
		return err
	}

	if liked == 1 {
		dislike, err = db.Prepare(`
			UPDATE postlikes
			SET liked = 0
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 0 {
		dislike, err = db.Prepare(`
			DELETE FROM postlikes
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	}
	_, err = dislike.Exec(
		uid,
		postid,
	)
	if err != nil {
		return err
	}

	return err
}

func LikeComment(uid int, postid int) error {
	like, err := db.Prepare(`
		INSERT INTO commentlikes
		(userid, commentid, liked)
		VALUES (?, ?, 1)
	`)
	if err != nil {
		return err
	}

	liked := 2
	if err = db.QueryRow(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND postid = ?)
	`, uid, postid).Scan(&uid, &postid, &liked); err != nil {
		return err
	}

	if liked == 1 {
		like, err = db.Prepare(`
			UPDATE commentlikes
			SET liked = 1
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 0 {
		like, err = db.Prepare(`
			DELETE FROM commentlikes
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	}
	_, err = like.Exec(
		uid,
		postid,
	)
	if err != nil {
		return err
	}
	return err
}
func DislikeComment(uid int, postid int) error {
	dislike, err := db.Prepare(`
		INSERT INTO commentlikes
		(userid, postid, liked)
		VALUES (?, ?, 0);
	`)
	if err != nil {
		return err
	}
	liked := 2
	if err = db.QueryRow(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND postid = ?)
	`, uid, postid).Scan(&uid, &postid, &liked); err != nil {
		return err
	}
	if liked == 1 {
		dislike, err = db.Prepare(`
			UPDATE commentlikes
			SET liked = 0
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	} else if liked == 0 {
		dislike, err = db.Prepare(`
			DELET FROM commentlikes
			WHERE (userid = ? AND postid = ?)
		`)
		if err != nil {
			return err
		}
	}
	_, err = dislike.Exec(
		uid,
		postid,
	)
	if err != nil {
		return err
	}
	return err
}
func getPostLikesDislikes(post *models.Post) error {
	if err := db.QueryRow(`
		SELECT COUNT(*)
		FROM postlikes
		WHERE (posid = ? AND liked = 1)
	`, post.PostID).Scan(&post.Like); err != nil {
		return err
	}

	if err := db.QueryRow(`
		SELECT COUNT(*)
		FROM postlikes
		WHERE (postid = ? AND liked = 0)
	`, post.PostID).Scan(&post.Dislike); err != nil {
		return err
	}
	return nil
}
func getCommentLikesDislikes(comment *models.Comment) error {
	if err := db.QueryRow(`
		SELECT COUNT(*)
		FROM commentlikes
		WHERE (commentid = ? AND liked = 1)
	`, comment.CommentID).Scan(&comment.Like); err != nil {
		return err
	}

	if err = db.QueryRow(`
		SELECT COUNT(*)
		FROM commentlikes
		WHERE (commentid = ? AND liked = 0)
	`, comment.CommentID).Scan(&comment.Dislike); err != nil {
		return err
	}

	return err
}
func postLikedByUser(post *models.Post, uid int) error {
	row, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ? AND liked = 1)
	`, uid, post.PostID)
	defer row.Close()

	if err != nil {
		return err
	}

	post.Liked = row.Next()

	return err
}

func postDislikedByUser(post *models.Post, uid int) error {
	row, err := db.Query(`
		SELECT *
		FROM postlikes
		WHERE (userid = ? AND postid = ? AND liked = 0)
	`, uid, post.PostID)

	if err != nil {
		return err
	}

	post.Disliked = row.Next()

	return err
}

func commentLikedByUser(comment *models.Comment, uid int) error {
	row, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ? AND liked = 1)
	`, uid, comment.CommentID)
	defer row.Close()

	if err != nil {
		return err
	}

	comment.Liked = row.Next()

	return err
}

func commentDislikedByUser(comment *models.Comment, uid int) error {
	row, err := db.Query(`
		SELECT *
		FROM commentlikes
		WHERE (userid = ? AND commentid = ? AND liked = 0)
	`, uid, comment.CommentID)
	defer row.Close()

	if err != nil {
		return err
	}

	comment.Disliked = row.Next()

	return err
}
