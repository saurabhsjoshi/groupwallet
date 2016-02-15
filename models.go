package main

import "time"

// User model in DB
type User struct {
	ID int64			`json:"id"`
	Name string			`json:"name"`
	Email string		`json:"email"`
	Password string		`json:"-"`
	Token string		`json:"token"`
	CreatedOn time.Time	`json:"createdOn"`
}

// Slice of users
type Users [] User

//Item model in DB
type Item struct {
	ID int64			`json:"id"`
	Owner int			`json:"owner"`
	Place string		`json:"place"`
	Extra string		`json:"extra"`
	Price float32		`json:"price"`
	CreatedOn time.Time	`json:"createdOn"`
}
// Slice of items
type Items[] Item

type Group struct {
	ID int64			`json:"id"`
	Title string		`json:"title"`
	Admin int64			`json:"admin"`
	GroupMembers Users	`json:"users"`
	GroupItems Items    `json:"items"`
}

//Slice of groups
type Groups[] Group

