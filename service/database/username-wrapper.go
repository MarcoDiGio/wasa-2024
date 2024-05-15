package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) ChangeUsername(user User) error {
	_, err := db.c.Exec("UPDATE INTO users VALUES (?)", user.ID)
	if err != nil {
		return err
	}
	return nil
}
