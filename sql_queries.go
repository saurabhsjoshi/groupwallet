package main

const (
	QUERY_INSERT_USER = `	INSERT INTO users (name, email, password, token, created_on)
		VALUES (?,?,?,?,NOW())
	`

	QUERY_INSERT_ITEM = `	INSERT INTO items (title, owner, place, extra, price, created_on)
		VALUES (?,?,?,?,?,NOW())
	`

	QUERY_INSERT_GROUP = `	INSERT INTO groups (title, admin, created_on)
		VALUES (?,?, NOW())
	`

	QUERY_INSERT_USER_INTO_GROUP = `	INSERT INTO ugr (u_id, g_id)
		VALUES (?,?)
	`

	QUERY_GET_ALL_USERS = "SELECT id, name, email, token, created_on FROM users"

	QUERY_GET_USER_BY_ID = "SELECT id, name, email, token FROM users WHERE id = ?"
)
