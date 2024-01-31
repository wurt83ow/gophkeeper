package models

import "time"

// Key is an alias for string and represents a key used in various contexts.
type Key string

// Response describes the server's response.
type Response struct {
	Result string `json:"result"`
}

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Deleted  bool   `db:"deleted"`
}

type UserCredential struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
	MetaInfo  string    `db:"meta_info"`
	Deleted   bool      `db:"deleted"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreditCardData struct {
	ID             int       `db:"id"`
	UserID         int       `db:"user_id"`
	CardNumber     string    `db:"card_number"`
	ExpirationDate string    `db:"expiration_date"`
	CVV            int       `db:"cvv"`
	MetaInfo       string    `db:"meta_info"`
	Deleted        bool      `db:"deleted"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type TextData struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Data      string    `db:"data"`
	MetaInfo  string    `db:"meta_info"`
	Deleted   bool      `db:"deleted"`
	UpdatedAt time.Time `db:"updated_at"`
}

type FileData struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Path      string    `db:"path"`
	MetaInfo  string    `db:"meta_info"`
	Deleted   bool      `db:"deleted"`
	UpdatedAt time.Time `db:"updated_at"`
}
