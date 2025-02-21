package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	items     = make(map[string]Item)
	idCounter = 1
	mutex     = &sync.Mutex{} // Mutex to handle concurrent access to the idCounter and items map
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createItem(w http.ResponseWriter, r *http.Request) {
	mutex.Lock() // Lock to ensure thread-safe access to idCounter
	defer mutex.Unlock()

	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.ID = strconv.Itoa(idCounter)
	idCounter++

	items[item.ID] = item
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	item, found := items[id]
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, found := items[id]
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	item.Name = updatedItem.Name
	item.Age = updatedItem.Age
	items[id] = item
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, found := items[id]
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(items, id)
	w.WriteHeader(http.StatusNoContent)
}

func main1() {
	http.HandleFunc("/create", createItem)
	http.HandleFunc("/get", getItem)
	http.HandleFunc("/getall", getAllItems)
	http.HandleFunc("/update", updateItem)
	http.HandleFunc("/delete", deleteItem)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// with id is uuid
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/google/uuid"
// )

// var items = make(map[string]Item)

// type Item struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func createItem(w http.ResponseWriter, r *http.Request) {
// 	var item Item
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&item); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	item.ID = uuid.New().String()
// 	items[item.ID] = item
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(item)
// }

// func getItem(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	item, found := items[id]
// 	if !found {
// 		http.Error(w, "Item not found", http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(item)
// }

// func updateItem(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	var updatedItem Item
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&updatedItem); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	item, found := items[id]
// 	if !found {
// 		http.Error(w, "Item not found", http.StatusNotFound)
// 		return
// 	}

// 	item.Name = updatedItem.Name
// 	item.Age = updatedItem.Age
// 	items[id] = item
// 	json.NewEncoder(w).Encode(item)
// }

// func deleteItem(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	_, found := items[id]
// 	if !found {
// 		http.Error(w, "Item not found", http.StatusNotFound)
// 		return
// 	}
// 	delete(items, id)
// 	w.WriteHeader(http.StatusNoContent)
// }

// func main() {
// 	http.HandleFunc("/create", createItem)
// 	http.HandleFunc("/get", getItem)
// 	http.HandleFunc("/update", updateItem)
// 	http.HandleFunc("/delete", deleteItem)

// 	fmt.Println("Server running on port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

/////////////////////////////////////////////////////////////////////////////////////////////////////////

// curl request

// curl -X POST http://localhost:8080/create \
// -H "Content-Type: application/json" \
// -d '{"name": "John", "age": 30}'

// curl "http://localhost:8080/get?id=1"

// curl -X PUT http://localhost:8080/update?id=1 \
// -H "Content-Type: application/json" \
// -d '{"name": "John Updated", "age": 31}'

// curl -X DELETE "http://localhost:8080/delete?id=1"
