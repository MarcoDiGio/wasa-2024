package database

func (db *appdbimpl) GetUserProfile() ([]UserProfile, error) {
	var ret []UserProfile
	rows, err := db.c.Query("SELECT id, followers, following FROM UserProfile")
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var f UserProfile
		err = rows.Scan(&f.ID, &f.Followers, &f.Following, &f.Photos)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}

	return ret, nil
}
