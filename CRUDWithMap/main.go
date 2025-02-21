package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[string]user)

func createUser(w http.ResponseWriter, r *http.Request) {
	var user user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	users[user.ID] = user
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user user
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	value, found := users[id]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	value.Name = user.Name
	value.Age = user.Age
	users[user.ID] = value
	// Respond with updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func main() {
	// http.Handle("/create", http.HandlerFunc(createUser)) //Correct use of HandlerFunc
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/getAll", getUsers)
	http.HandleFunc("/update", UpdateUser)
	fmt.Println("Server running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
