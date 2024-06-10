package database

func (db *appdbimpl) AddLike(user User, photoId string) error {
	_, err := db.c.Exec("INSERT INTO likes(liker_id, photo_id) VALUES (?, ?)", user.ID, photoId)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetLikesByPhoto(photoId string) ([]User, error) {
	users := make([]User, 0)
	rows, err := db.c.Query("SELECT liker_id FROM likes WHERE photo_id=?", photoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var like Like
		if err := rows.Scan(&like.Liker_ID); err != nil {
			return nil, err
		}
		// since there is a foreign key costraint, there is no need
		// to check if user exists
		users = append(users, User{ID: like.Liker_ID})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (db *appdbimpl) RemoveLike(user User, photoId string) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE liker_id=? AND photo_id=?", user.ID, photoId)
	if err != nil {
		return err
	}
	return nil
}
