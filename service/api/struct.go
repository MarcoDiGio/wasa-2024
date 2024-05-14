package api

import "wasa-2024-2024851/service/database"

type User struct {
	ID string
}

func (u *User) toDatabase() database.User {
	return database.User{
		ID: u.ID,
	}
}

func (u *User) isValid() bool {
	//@TODO
	return true
}
