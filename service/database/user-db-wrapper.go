package database

func (db *appdbimpl) PostUser(u User) error {
	var exists bool
	var err error
	exists, err = db.CheckUser(u)

	if err != nil {
		return err
	}

	if exists {
		return nil
	} else {
		_, err := db.c.Exec("INSERT INTO users VALUES(?)", u.ID)
		if err != nil {
			return err
		}
		return err
	}
}

func (db *appdbimpl) CheckUser(u User) (bool, error) {
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id=?)", u.ID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (db *appdbimpl) GetAllUsers() ([]User, error) {
	var users = make([]User, 0)
	rows, err := db.c.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *appdbimpl) ChangeUsername(user User) error {
	_, err := db.c.Exec("UPDATE INTO users VALUES (?)", user.ID)
	if err != nil {
		return err
	}
	return nil
}
