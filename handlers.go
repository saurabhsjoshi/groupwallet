package main
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"net/http"
	"os"
	"encoding/json"
	"strconv"
)

/* TODO: Clean up db code and manage errors gracefully*/
/* Handlers for different endpoints */


func addHeaders(w http.ResponseWriter) http.ResponseWriter{
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return w
}

func connectToDb() (*sql.DB, error) {
	connectionString := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_URL") + ":3306)/" + os.Getenv("DB_NAME")
	return sql.Open("mysql", connectionString)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDb()

	if(err!= nil) {
		json.NewEncoder(w).Encode(NewDbErrorStatus())
		panic(err)
	}

	owner_id,err := strconv.ParseInt(r.FormValue("owner"),10, 64)
	if(err != nil){
		panic(err)
	}

	price, err := strconv.ParseFloat(r.FormValue("price"),64)
	if(err!=nil){
		panic(err)
	}

	item := Item{
		Title: r.FormValue("title"),
		Owner: owner_id,
		Place: r.FormValue("place"),
		Extra: r.FormValue("extra"),
		Price: price,
	}

	stmt, err := db.Prepare(QUERY_INSERT_ITEM)
	if(err != nil){
		panic(err)
	}

	res, err := stmt.Exec(
		item.Title,
		item.Owner,
		item.Place,
		item.Extra,
		item.Price,
	)

	if(err != nil){
		panic(err)
	}

	item.ID,_ = res.LastInsertId()

	json.NewEncoder(w).Encode(
			struct{
			StatusMessage
			Item
		} {
			NewSuccessStatus(),
			item,
		})
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	/* TODO: Add auth and hash password */
	db, err := connectToDb()

	if err != nil {
		json.NewEncoder(w).Encode(NewDbErrorStatus())
		panic(err)
	}

	admin, err := strconv.ParseInt(r.FormValue("adminId"), 10, 64)
	if(err!=nil){
		panic(err)
	}
	group := Group{
		Title: r.FormValue("name"),
		Admin: admin,
	}

	stmt, err := db.Prepare(QUERY_INSERT_GROUP)

	if err != nil{
		panic (err)
	}

	res, err := stmt.Exec(
		group.Title,
		group.Admin,
	)
	if err!= nil{
		panic(err)
	}

	group.ID, _ = res.LastInsertId()

	json.NewEncoder(w).Encode(
			struct{
			StatusMessage
			Group
		} {
			NewSuccessStatus(),
			group,
		})


}

/* Add a new user to the database */
func RegisterUser (w http.ResponseWriter, r *http.Request) {
	/* TODO: Add auth and hash password */
	db, err := connectToDb()

	if err != nil {
		json.NewEncoder(w).Encode(NewDbErrorStatus())
		panic(err)
	}

	user := User{
		Name: r.FormValue("name"),
		Email:r.FormValue("email"),
		Password:r.FormValue("password"),
		Token:"",
	}

	stmt, err := db.Prepare(QUERY_INSERT_USER)

	if err != nil{
		panic(err)
	}

	res, err := stmt.Exec(
		user.Name,
		user.Email,
		user.Password,
		user.Token,
	)
	if err!=nil {
		panic(err)
	}
	user.ID,_ = res.LastInsertId()

	json.NewEncoder(w).Encode(
			struct{
			StatusMessage
			User
		} {
			NewSuccessStatus(),
			user,
		})
}

/* Just for testing db */
func GetAllUsers (w http.ResponseWriter, r *http.Request){

	w = addHeaders(w)
	db, err := connectToDb()

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


