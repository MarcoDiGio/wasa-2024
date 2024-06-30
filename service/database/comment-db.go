package database

func (db *appdbimpl) AddComment(photoId string, user User, comment Comment) (int64, error) {
	res, err := db.c.Exec("INSERT INTO comments(photo_id, user_id, comment) VALUES (?, ?, ?)", photoId, user.ID, comment.Comment)
	if err != nil {
		return -999, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -999, err
	}
	return id, nil
}

func (db *appdbimpl) CheckCommentAuthor(commentId string, photoId string, user User) (bool, error) {
	var isCommenter bool
	err := db.c.QueryRow("SELECT COUNT(*) FROM comments WHERE comment_id=? AND photo_id=? AND user_id=?", commentId, photoId, user.ID).Scan(&isCommenter)
	if err != nil {
		return false, err
	}
	return isCommenter, nil
}

func (db *appdbimpl) RemoveComment(commentId string) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE comment_id=?", commentId)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetCommentsByPhoto(photoId string) ([]Comment, error) {
	var comments = make([]Comment, 0)
	rows, err := db.c.Query("SELECT * FROM comments WHERE photo_id=?", photoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.Comment_ID, &comment.Photo_ID, &comment.User_ID, &comment.Comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
