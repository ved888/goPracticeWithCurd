package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var users = make(map[string]User)

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid user request", http.StatusBadRequest)
		return
	}
	users[user.ID] = user
	w.Header().Set("contant-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contant-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	user, exist := users[id]
	if !exist {
		http.Error(w, "user is not exist", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}

func update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid user request", http.StatusBadRequest)
		return
	}
	u, exist := users[id]
	if !exist {
		http.Error(w, "user is not found", http.StatusInternalServerError)
		return
	}
	u.Name = user.Name
	u.Add = user.Add
	users[id] = u

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	_, found := users[id]
	if !found {
		http.Error(w, "user is not found", http.StatusNotFound)
		return
	}
	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/getAll", getAll).Methods("GET")
	r.HandleFunc("/getById/{id}", getById).Methods("GET")
	r.HandleFunc("/update/{id}", update).Methods("PUT")
	r.HandleFunc("/Delete/{id}", deleteUser).Methods("DELETE")

	fmt.Println("server start on poart 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
