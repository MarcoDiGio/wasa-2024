package database

func (db *appdbimpl) AddFollower(follower User, followed User) error {
	var err error
	_, err = db.c.Exec("INSERT INTO followers(follower, followed) VALUES (?, ?)", follower.ID, followed.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetFollower(user User) ([]User, error) {
	var followers = make([]User, 0)
	rows, err := db.c.Query("SELECT follower FROM followers WHERE followed=?", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var follower User
		err := rows.Scan(&follower.ID)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return followers, nil
}

func (db *appdbimpl) GetFollowing(user User) ([]User, error) {
	var followings = make([]User, 0)
	rows, err := db.c.Query("SELECT followed FROM followers WHERE follower=?", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var following User
		err := rows.Scan(&following.ID)
		if err != nil {
			return nil, err
		}
		followings = append(followings, following)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return followings, nil
}

func (db *appdbimpl) RemoveFollower(follower User, followed User) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE follower = ? AND followed = ?", follower.ID, followed.ID)
	if err != nil {
		return err
	}
	return nil
}
