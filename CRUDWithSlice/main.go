package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func getByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for _, v := range users {
		if v.Id == id {
			w.Header().Set("content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)

}

func upadateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil || userID < 0 || userID >= len(users) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "request body is not valid", http.StatusBadRequest)
		return
	}

	users[userID].Name = user.Name
	users[userID].Add = user.Add

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// fix above id problem
func updateUserById(w http.ResponseWriter, r *http.Request) {
	var user User
	id := r.URL.Query().Get("id")

	// Find user by ID in slice
	index := -1
	for i, u := range users {
		if u.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode request body into user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Request body is not valid", http.StatusBadRequest)
		return
	}

	// Update user details in slice
	users[index].Name = user.Name
	users[index].Add = user.Add

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users[index])
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(id)
	if err != nil || userID < 0 || userID > len(users) {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	users = append(users[:userID], users[userID+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

// /////////////////////////OR/////////////////////.
func deleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Create a new slice excluding the user with the given ID
	newUsers := []User{}
	for _, user := range users {
		if user.Id != id { // Keep all except matching ID
			newUsers = append(newUsers, user)
		}
	}

	// Update the original users slice
	users = newUsers

	// Send response
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/getAll", getAllUser)
	http.HandleFunc("/get", getByID)
	http.HandleFunc("/update", upadateUser)
	http.HandleFunc("/updateByID", updateUserById)
	http.HandleFunc("/delete", deleteUser)
	http.HandleFunc("/deleteByID", deleteUserByID)

	fmt.Println("server start on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
