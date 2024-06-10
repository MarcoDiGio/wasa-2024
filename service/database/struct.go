package database

import "time"

type User struct {
	ID string `json:"user_id"`
}

type UserProfile struct {
	ID         string       `json:"user_id"`
	Followings []User       `json:"followings"`
	Followers  []User       `json:"followers"`
	Photos     []FinalPhoto `json:"photos"`
}

type Photo struct {
	Photo_ID  string    `json:"photo_ID"`
	Author_ID string    `json:"author_id"`
	Date      time.Time `json:"date"`
}

type FinalPhoto struct {
	Photo_ID  string    `json:"photo_ID"`
	Author_ID string    `json:"author_id"`
	Date      time.Time `json:"date"`
	Comments  []Comment `json:"comments"`
	Likes     []User    `json:"likes"`
}

type Comment struct {
	Comment_ID string `json:"comment_id"`
	Photo_ID   string `json:"photo_id"`
	User_ID    string `json:"user_id"`
	Comment    string `json:"comment"`
}

type Like struct {
	Liker_ID string `json:"liker_id"`
	Photo_ID string `json:"photo_id"`
}
