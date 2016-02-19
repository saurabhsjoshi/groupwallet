package main
import "time"

//Item model in DB
type Item struct {
	ID int64			`json:"id"`
	Title string        `json:"title"`
	Owner int64			`json:"owner"`
	Place string		`json:"place"`
	Extra string		`json:"extra"`
	Price float64		`json:"price"`
	CreatedOn time.Time	`json:"createdOn"`
}
// Slice of items
type Items[] Item
