package main


const (

	QUERY_INSERT_USER =
	`	INSERT INTO users (name, email, password, token, created_on)
		VALUES (?,?,?,?,NOW())
	`

	QUERY_INSERT_ITEM =
	`	INSERT INTO items (title, owner, place, extra, price, created_on)
		VALUES (?,?,?,?,?,NOW())
	`

	QUERY_GET_ALL_USERS = "SELECT id, name, email, token, created_on FROM users"
)