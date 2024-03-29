package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
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

func (user *User) GetFromDb(id int64) error {
	db, err := ConnectToDb()

	if err != nil {
		return NewDbErrorStatus("Could not connect to db.", err)
	}

	err = db.QueryRow(QUERY_GET_USER_BY_ID, id).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Token, &user.CreatedOn)

	if err == sql.ErrNoRows {
		return NewNotFoundErrorStatus("User with id "+strconv.FormatInt(id, 10)+" not found.", err)
	} else if err != nil {
		return NewDbErrorStatus("Failed to get user from db.", err)
	}

	return nil
}

func (user *User) PutInDb() (int64, error) {
	db, err := ConnectToDb()

	if err != nil {
		return -1, NewDbErrorStatus("Could not connect to db.", err)
	}

	stmt, err := db.Prepare(QUERY_INSERT_USER)

	if err != nil {
		return -1, NewDbErrorStatus("Could not create statement.", err)
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

func (user *User) UpdateInDb() error {
	db, err := ConnectToDb()

	if err != nil {
		return NewDbErrorStatus("Could not connect to db.", err)
	}

	if stmt, err := db.Prepare(QUERY_UPDATE_USER); err != nil {
		return NewDbErrorStatus("Could not create statement.", err)
	} else {
		if _, err := stmt.Exec(
			user.Name,
			user.Email,
			user.Password,
			user.Token,
		); err != nil {
			return NewDbErrorStatus("Could not update user.", err)
		}
	}

	return nil
}

func (user *User) DeleteFromDb() error {
	db, err := ConnectToDb()

	if err != nil {
		return NewDbErrorStatus("Could not connect to db.", err)
	}

	if _, err := db.Exec(QUERY_DELETE_USER_BY_ID, user.ID); err != nil {
		return NewDbErrorStatus("Could not delete user.", err)
	}

	return nil
}
