package main
import "time"

type Group struct {
	ID int64			`json:"id"`
	Title string		`json:"title"`
	Admin int64			`json:"admin"`
	GroupMembers Users	`json:"users"`
	GroupItems Items    `json:"items"`
	CreatedOn time.Time `json:"createdOn"`
}

//Slice of groups
type Groups[] Group
