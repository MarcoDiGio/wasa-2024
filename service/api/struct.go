package api

import (
	"time"
	"wasa-2024-2024851/service/database"
)

type User struct {
	ID string `json:"user_id"`
}

type UserProfile struct {
	ID         string                `json:"user_id"`
	Followings []database.User       `json:"followings"`
	Followers  []database.User       `json:"followers"`
	Photos     []database.FinalPhoto `json:"photos"`
}

type Photo struct {
	Photo_ID  string    `json:"photo_ID"`
	Author_ID string    `json:"author_id"`
	Date      time.Time `json:"date"`
}

type FinalPhoto struct {
	Photo_ID  string             `json:"photo_ID"`
	Author_ID string             `json:"author_id"`
	Date      time.Time          `json:"date"`
	Comments  []database.Comment `json:"comments"`
	Likes     []database.User    `json:"likes"`
}

type Comment struct {
	Comment_ID string `json:"comment_id"`
	Photo_ID   string `json:"photo_id"`
	User_ID    string `json:"user_id"`
	Comment    string `json:"comment"`
}

func (user User) toDatabase() database.User {
	return database.User{
		ID: user.ID,
	}
}

func (photo Photo) toDatabase() database.Photo {
	return database.Photo{
		Photo_ID:  photo.Photo_ID,
		Author_ID: photo.Author_ID,
		Date:      photo.Date,
	}
}

func (comment Comment) toDatabase() database.Comment {
	return database.Comment{
		Comment_ID: comment.Comment_ID,
		Photo_ID:   comment.Photo_ID,
		User_ID:    comment.User_ID,
		Comment:    comment.Comment,
	}
}

func (profile UserProfile) toDatabase() database.UserProfile {
	return database.UserProfile{
		ID:         profile.ID,
		Followings: profile.Followings,
		Followers:  profile.Followers,
		Photos:     profile.Photos,
	}
}
