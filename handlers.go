package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"strconv"
)

/* TODO: Clean up db code and manage errors gracefully*/
/* Handlers for different endpoints */

func connectToDb() (*sql.DB, error) {
	connectionString := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_URL") + ":3306)/" + os.Getenv("DB_NAME")
	return sql.Open("mysql", connectionString)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDb()

	if err != nil {
		panic(NewDbErrorStatus("Could not connect to db.", err))
	}

	var item Item

	if err := item.UnmarshallHTTP(r); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare(QUERY_INSERT_ITEM)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(
		item.Title,
		item.Owner,
		item.Place,
		item.Extra,
		item.Price,
	)

	if err != nil {
		panic(err)
	}

	item.ID, _ = res.LastInsertId()

	json.NewEncoder(w).Encode(
		struct {
			StatusMessage
			Item
		}{
			NewSuccessStatus(),
			item,
		})
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	/* TODO: Add auth and hash password */
	db, err := connectToDb()

	if err != nil {
		panic(NewDbErrorStatus("Could not connect to db.", err))
	}

	admin, err := strconv.ParseInt(r.FormValue("adminId"), 10, 64)
	if err != nil {
		panic(err)
	}
	group := Group{
		Title: r.FormValue("name"),
		Admin: admin,
	}

	stmt, err := db.Prepare(QUERY_INSERT_GROUP)

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(
		group.Title,
		group.Admin,
	)
	if err != nil {
		panic(err)
	}

	group.ID, _ = res.LastInsertId()

	json.NewEncoder(w).Encode(
		struct {
			StatusMessage
			Group
		}{
			NewSuccessStatus(),
			group,
		})

}

/* Get user by id */
func GetUserById(w http.ResponseWriter, r *http.Request) {
	var u User
	id, err := strconv.ParseInt(r.FormValue("userId"), 10, 64)

	if err != nil {
		panic(NewUnknownErrorStatus())
	}

	err = u.GetFromDb(id)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(
		struct {
			StatusMessage
			User `json:"User"`
		}{
			NewSuccessStatus(),
			u,
		})
}

/* Add a new user to the database */
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	/* TODO: Add auth and hash password */

	var u User

	if err := u.UnmarshallHTTP(r); err != nil {
		panic(err)
	}

	if _, err := u.PutInDb(); err != nil {
		panic(err)
	} else {
		json.NewEncoder(w).Encode(
			struct {
				StatusMessage
				User
			}{
				NewSuccessStatus(),
				u,
			})
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if user_id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err != nil {
		panic(NewUnmarshallErrorStatus("Could not parse admin id.", err))
	} else {
		u.ID = user_id
	}

	if err := u.UnmarshallHTTP(r); err != nil {
		panic(err)
	}

	if err := u.DeleteFromDb(); err != nil {
		panic(err)
	} else {
		json.NewEncoder(w).Encode(NewSuccessStatus())
	}
}

/* Just for testing db */
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDb()

	if err != nil {
		panic(NewDbErrorStatus("Could not connect to db.", err))
	}

	rows, err := db.Query(QUERY_GET_ALL_USERS)

	if err != nil {
		panic(NewDbErrorStatus("Could not get users.", err))
	}

	users := Users{}

	for rows.Next() {
		var user User
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Token,
			&user.CreatedOn,
		)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(
		struct {
			StatusMessage
			Users
		}{
			NewSuccessStatus(),
			users,
		})

	db.Close()
}
