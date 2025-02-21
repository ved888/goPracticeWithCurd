package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var users []User

func create(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/getAll", getAllUser)

	fmt.Println("server start on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
