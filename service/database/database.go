/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	ChangeUsername(user User) error
	CheckUser(user User) (bool, error)
	PostUser(user User) error
	GetAllUsers() ([]User, error)
	Ping() error
}

type User struct {
	ID string `json:"user_id"`
}

type UserProfile struct {
	ID        uint64
	Following uint64
	Followers uint64
	Photos    []Photo
}

type Photo struct {
	ID string
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func createDatabase(db *sql.DB) error {
	// @TODO: Think more about the DB structure
	tables := [6]string{
		`CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(16) NOT NULL PRIMARY KEY
		);`,
		`CREATE TABLE IF NOT EXISTS photos (
			photo_id INTEGER PRIMARY KEY AUTOINCREMENT,
			date DATETIME NOT NULL,
			author_id VARCHAR(16) NOT NULL,
			FOREIGN KEY(author_id) REFERENCES users (id_user) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			photo_id INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS likes (
			like_id INTEGER PRIMARY KEY AUTOINCREMENT,
			photo_id INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS banned_users (
			banner VARCHAR(16) NOT NULL,
			banned VARCHAR(16) NOT NULL,
			PRIMARY KEY(banner, banned),
			FOREIGN KEY(banner) REFERENCES users (id_user) ON DELETE CASCADE,
			FOREIGN KEY(banned) REFERENCES users (id_user) ON DELETE CASCADE,
			CONSTRAINT CHK_ForeignKeysDifferent CHECK (banner <> banned)
		);`,
		`CREATE TABLE IF NOT EXISTS follower (
			follower VARCHAR(16) NOT NULL,
			followed VARCHAR(16) NOT NULL,
			PRIMARY KEY(follower, followed)
			FOREIGN KEY(follower) REFERENCES users (id_user) ON DELETE CASCADE,
			FOREIGN KEY(followed) REFERENCES users (id_user) ON DELETE CASCADE,
			CONSTRAINT CHK_ForeignKeysDifferent CHECK (follower <> followed)
		);`,
	}

	for i := 0; i < len(tables); i++ {
		_, err := db.Exec(tables[i])
		if err != nil {
			return fmt.Errorf("error creating database structure: %w", err)
		}
	}
	return nil
}
