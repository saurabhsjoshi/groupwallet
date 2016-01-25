package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request){
	user := User{
		ID: 0,
		Name: "Saurabh Joshi",
		Email: "abc@xyz.com",
		Password: "saltedpassword",
	}
	if err := json.NewEncoder(w).Encode(user); err != nil{
		panic(err)
	}
}