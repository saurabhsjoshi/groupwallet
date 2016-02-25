package main

import (
	"errors"
	"net/http"
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
		return TNewDbErrorStatus("Could not connect to db.", err)
	}

	err = db.QueryRow(QUERY_GET_USER_BY_ID, id).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Token, &user.CreatedOn)

	if err != nil {
		return TNewDbErrorStatus("Failed to get user from db.", err)
	}

	return nil
}

func (user *User) PutInDb() (int64, error) {
	db, err := ConnectToDb()

	if err != nil {
		return -1, TNewDbErrorStatus("Could not connect to db.", err)
	}
	stmt, err := db.Prepare(QUERY_INSERT_USER)
	if err != nil {
		return -1, TNewDbErrorStatus("Could not create statement.", err)
	}
	user.CreatedOn = time.Now()

	res, err := stmt.Exec(
		user.Name,
		user.Email,
		user.Password,
		user.Token,
		user.CreatedOn,
	)
	if err != nil {
		panic(err)
	}
	user.ID, _ = res.LastInsertId()

	return user.ID, nil
}

func (user *User) UnmarshallHTTP(r *http.Request) error {

	defer func() error {
		if r := recover(); r != nil {
			return NewUnmarshallErrorStatus("Incorrect user values.", errors.New("Could not unmarshall."))
		} else {
			return nil
		}
	}()

	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	user.Token = ""

	return nil
}
