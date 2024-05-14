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
