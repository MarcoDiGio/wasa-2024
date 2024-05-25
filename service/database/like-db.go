package database

func (db *appdbimpl) AddLike(user User, photoId string) error {
	_, err := db.c.Exec("INSERT INTO likes(liker_id, photo_id) VALUES (?, ?)", user.ID, photoId)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveLike(user User, photoId string) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE liker_id=? AND photo_id=?", user.ID, photoId)
	if err != nil {
		return err
	}
	return nil
}
