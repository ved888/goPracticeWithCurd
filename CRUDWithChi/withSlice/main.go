package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var Users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	Users = append(Users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "create user"})
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userList []User

	for _, user := range Users {
		userList = append(userList, User{Id: user.Id, Name: user.Name, Add: user.Add})
	}
	json.NewEncoder(w).Encode(userList)
}

func getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")

	for _, value := range Users {
		if value.Id == id {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	index := -1
	for i, values := range Users {
		if values.Id == id {
			index = i
		}
	}
	Users[index].Name = user.Name
	Users[index].Add = user.Add
	w.Header().Set("contant-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "updated user"})

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newUser = []User{}

	for _, user := range Users {
		if user.Id != id {
			newUser = append(newUser, user)
		}
	}
	Users = newUser
	json.NewEncoder(w).Encode(map[string]string{"message": "delete user"})

}
func main() {
	r := chi.NewRouter()

	r.Post("/create", createUser)
	r.Get("/getAll", getAll)
	r.Get("/get", getById)
	r.Put("/update/{id}", updateUser)
	r.Delete("/delete/{id}", deleteUser)

	fmt.Println("server start on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", r))
}
