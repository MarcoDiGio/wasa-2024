package database

func (db *appdbimpl) AddPhoto(photo Photo) (int64, error) {
	res, err := db.c.Exec("INSERT INTO photos(author_id, date) VALUES (?, ?)", photo.Author_ID, photo.Date)
	if err != nil {
		return -999, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -999, err
	}

	return id, nil
}

func (db *appdbimpl) RemovePhoto(photoId string) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE photo_id=?", photoId)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CheckPhoto(photoId string) (bool, error) {
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM photos WHERE photoId=?)", photoId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (db *appdbimpl) GetAllUserPhotos(user User) ([]Photo, error) {
	var photos = make([]Photo, 0)
	rows, err := db.c.Query("SELECT * FROM photos WHERE author_id=? ORDER BY date DESC", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.Photo_ID, &photo.Date, &photo.Author_ID); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}

func (db *appdbimpl) GetAllFollowingPhotos(user User) ([]Photo, error) {
	var photos = make([]Photo, 0)
	rows, err := db.c.Query("SELECT * FROM photos WHERE author_id IN (SELECT followed FROM followers WHERE follower=?) ORDER BY date DESC", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.Photo_ID, &photo.Date, &photo.Author_ID); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}
