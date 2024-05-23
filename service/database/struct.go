package database

import "time"

type User struct {
	ID string `json:"user_id"`
}

type UserProfile struct {
	ID         string  `json:"user_id"`
	Followings []User  `json:"followings"`
	Followers  []User  `json:"followers"`
	Photos     []Photo `json:"photos"`
}

type Photo struct {
	Photo_ID  string    `json:"photo_ID"`
	Author_ID string    `json:"author_id"`
	Date      time.Time `json:"date"`
}

type Comment struct {
	Comment_ID string `json:"comment_id"`
	Photo_ID   string `json:"photo_id"`
	User_ID    string `json:"user_id"`
	Comment    string `json:"content"`
}
