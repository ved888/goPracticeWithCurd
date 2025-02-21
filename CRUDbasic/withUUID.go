package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Object struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quentity int    `json:"quentity"`
}

var objects = make(map[string]Object)

func createObject(w http.ResponseWriter, r *http.Request) {
	var object Object

	err := json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		return
	}
	object.Id = uuid.New().String()
	objects[object.Id] = object
	w.Header().Set("contant-type", "application/json")
	json.NewEncoder(w).Encode(object)
}

func getObject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	object, fond := objects[id]
	if !fond {
		http.Error(w, "object not found", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(object)
}

func getAllObject(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(objects)
}

func main() {
	http.HandleFunc("/create", createObject)
	http.HandleFunc("/get", getObject)
	http.HandleFunc("/getAll", getAllObject)
	fmt.Println("satrt server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
