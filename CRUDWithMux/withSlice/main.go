package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	users = append(users, user)
	w.Header().Set("contant-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("contant-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for _, user := range users {
		if user.Id == id {
			w.Header().Set("contant-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	index := -1

	for i, u := range users {
		if u.Id == id {
			index = i
		}
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	users[index].Name = user.Name
	users[index].Add = user.Add

	w.Header().Set("contant-Type", "application/json")
	json.NewEncoder(w).Encode(users[index])
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	newUser := []User{}

	for _, user := range users {
		if user.Id != id {
			newUser = append(newUser, user)
		}
	}
	users = newUser
	w.WriteHeader(http.StatusNoContent)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/getAll", getAll).Methods("get")
	r.HandleFunc("/get", getById).Methods("get")
	r.HandleFunc("/update", updateUser).Methods("PUT")
	r.HandleFunc("/delete", deleteUser).Methods("delete")

	log.Println("Server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))

}
