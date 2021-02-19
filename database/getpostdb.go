package database

import (
	"forum/models"
)

func GetPosts(uid int) ([]models.Post, error) {
	var posts []models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
		ORDER BY date_time DESC
	`)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.ImageExist, &post.DateTime, &post.TimeString)
		err = getPostLikesDislikes(&post)
		if err != nil {
			return posts, err
		}
		if uid != 0 {
			err = postLikedByUser(&post, uid)
			if err != nil {
				return posts, err
			}
			err = postDislikedByUser(&post, uid)
			if err != nil {
				return posts, err
			}
		}
		err = getPostCommentCount(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetPostByID(postid int, uid int) (models.Post, error) {
	var post models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE postid = ?
	`, postid)
	defer row.Close()

	if err != nil {
		return post, err
	}

	for row.Next() {
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.ImageExist, &post.DateTime, &post.TimeString)
	}

	err = getPostCategories(&post)
	if err != nil {
		return post, err
	}

	err = getPostLikesDislikes(&post)
	if err != nil {
		return post, err
	}

	if uid != 0 {
		err = postLikedByUser(&post, uid)
		if err != nil {
			return post, err
		}
		err = postDislikedByUser(&post, uid)
		if err != nil {
			return post, err
		}
	}
	err = getPostCommentCount(&post)
	if err != nil {
		return post, err
	}

	return post, err
}

func GetPostsByCategory(category string, userid int) ([]models.Post, error) {
	var posts []models.Post

	row, err := db.Query(`
		SELECT posts.*
		FROM posts
		INNER JOIN postcategories
		ON postcategories.postid = posts.postid
		INNER JOIN categories
		ON categories.categoryid = postcategories.categoryid
		WHERE categories.categoryname = ?
	`, category)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.ImageExist, &post.DateTime, &post.TimeString)
		err = getPostLikesDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
			err = postDislikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		err = getPostCommentCount(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, err
}

func GetPostsByUserID(userid int) ([]models.Post, error) {
	var posts []models.Post

	row, err := db.Query(`
		SELECT *
		FROM posts
		WHERE userid = ?
	`, userid)
	defer row.Close()

	if err != nil {
		return posts, err
	}

	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.ImageExist, &post.DateTime, &post.TimeString)
		err = getPostLikesDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
			err = postDislikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		err = getPostCommentCount(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetLikedPostsByUserID(userid int) ([]models.Post, error) {
	var posts []models.Post

	row, err := db.Query(`
		SELECT posts.*
		FROM posts
		INNER JOIN postlikes
		ON posts.postid = postlikes.postid
		WHERE postlikes.userid = ? AND postlikes.liked = 1
	`, userid)
	defer row.Close()

	if err != nil {
		return posts, err
	}
	for row.Next() {
		var post models.Post
		row.Scan(&post.PostID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.ImageExist, &post.DateTime, &post.TimeString)
		err = getPostLikesDislikes(&post)
		if err != nil {
			return posts, err
		}
		if userid != 0 {
			err = postLikedByUser(&post, userid)
			if err != nil {
				return posts, err
			}
		}
		err = getPostCommentCount(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
