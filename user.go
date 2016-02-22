package main

import (
	"log"
	"time"
)

// User model in DB
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Token     string    `json:"token"`
	CreatedOn time.Time `json:"createdOn"`
}

// Slice of users
type Users []User

func (user *User) GetFromDb(id int64) error {
	db, err := ConnectToDb()

	if err != nil {
		return TNewDbErrorStatus()
	}

	err = db.QueryRow(QUERY_GET_USER_BY_ID, id).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Token)

	if err != nil {
		log.Print("Failed to get user by ID ", err)
		return TNewDbErrorStatus()
	}

	return nil
}
