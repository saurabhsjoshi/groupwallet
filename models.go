package main

import "time"

// User model in DB
type User struct {
	ID int    			`json:"id"`
	Name string     	`json:"name"`
	Email string    	`json:"email"`
	Password string 	`json:"-"`
	CreatedOn time.Time `json:"createdOn"`
}

// Slice of users
type Users [] User

//Item model in DB
type Item struct {
	ID int            	`json:"id"`
	Owner int           `json:"owner"`
	Place string        `json:"place"`
	Extra string        `json:"extra"`
	Price float32       `json:"price"`
	CreatedOn time.Time `json:"createdOn"`
}

// Slice of items
type Items[] Item