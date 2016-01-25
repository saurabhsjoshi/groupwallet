package main
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"net/http"
	"os"
	"encoding/json"
)

/* Handlers for different endpoints */


/* Add a new user to the database */
func RegisterUser (w http.ResponseWriter, r *http.Request) {

}

/* Just for testing db */
func GetAllUsers (w http.ResponseWriter, r *http.Request){
	connectionString := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_URL") + ":3306)/" + os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", connectionString)

	if(err != nil){
		json.NewEncoder(w).Encode(NewDbErrorStatus())
		panic(err)
	}

	rows, err := db.Query(QUERY_GET_ALL_USERS)
	if(err != nil){
		json.NewEncoder(w).Encode(NewDbErrorStatus())
		panic(err)
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
		users = append(users,user)
	}
	json.NewEncoder(w).Encode(
		struct{
			StatusMessage
			Users
		} {
			NewSuccessStatus(),
			users,
		})

	db.Close()

}


