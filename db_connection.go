package main

import (
	"database/sql"
	"os"
)

/* Returns an open connection to the database */
func ConnectToDb() (*sql.DB, error) {
	connectionString := os.Getenv("DB_USERNAME") + ":" +
		os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_URL") +
		":3306)/" +
		os.Getenv("DB_NAME") + "?&charset=utf8&parseTime=True"
	return sql.Open("mysql", connectionString)
}
