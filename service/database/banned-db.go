package database

func (db *appdbimpl) AddBan(banner User, banned User) error {
	var err error
	_, err = db.c.Exec("INSERT INTO banned_users(banner, banned) VALUES (?, ?)", banner.ID, banned.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) RemoveBan(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned_users WHERE banner = ? AND banned = ?", banner.ID, banned.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) CheckBan(banner User, banned User) (bool, error) {
	var isBanned bool
	err := db.c.QueryRow("SELECT COUNT(*) FROM banned_users WHERE banner=? AND banned=?", banner.ID, banned.ID).Scan(&isBanned)
	if err != nil {
		return false, err
	}
	return isBanned, nil
}
