package main

import "net/http"

type Entity interface  {

	/* Populates the struct with HTTP params */
	UnmarshallHTTP(*http.Request) error;

	/* Populates the struct from DB with ID*/
	GetFromDb(id int) error;

	/* Adds the entity to the DB */
	PutInDb() (int, error);

	/* Updates the db with new values from struct */
	UpdateInDb() error;

	/* Deletes entity from the db */
	DeleteFromDb() error;
}
