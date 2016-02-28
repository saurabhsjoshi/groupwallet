package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

//Item model in DB
type Item struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Owner     int64     `json:"owner"`
	Place     string    `json:"place"`
	Extra     string    `json:"extra"`
	Price     float64   `json:"price"`
	CreatedOn time.Time `json:"createdOn"`
}

// Slice of items
type Items []Item

func (item *Item) UnmarshallHTTP(r *http.Request) error {
	defer func() error {
		if r := recover(); r != nil {
			return NewUnmarshallErrorStatus("Incorrect user values.", errors.New("Could not unmarshall."))
		} else {
			return nil
		}
	}()

	owner_id, err := strconv.ParseInt(r.FormValue("owner"), 10, 64)
	if err != nil {
		return NewUnmarshallErrorStatus("Could not parse owner id.", err)
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		return NewUnmarshallErrorStatus("Could not parse price of item.", err)
	}
	item.Title = r.FormValue("title")
	item.Owner = owner_id
	item.Place = r.FormValue("place")
	item.Extra = r.FormValue("extra")
	item.Price = price

	return nil
}
