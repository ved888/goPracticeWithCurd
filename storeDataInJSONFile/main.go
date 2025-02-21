package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	items     = make(map[string]Item)
	idCounter = 1
	mutex     = &sync.Mutex{}
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func loadItemsFromFile() error {
	file, err := os.Open("items.json")
	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist, initialize empty items
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&items)
	if err != nil {
		return err
	}

	// Find the highest ID to continue from for the next item
	for id := range items {
		itemID, _ := strconv.Atoi(id)
		if itemID >= idCounter {
			idCounter = itemID + 1
		}
	}
	return nil
}

func saveItemsToFile() error {
	file, err := os.Create("items.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var item Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.ID = strconv.Itoa(idCounter)
	idCounter++

	items[item.ID] = item

	if err := saveItemsToFile(); err != nil {
		http.Error(w, "Failed to save item", http.StatusInternalServerError)
		return
	}

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

func updateItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedItem Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
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

	if err := saveItemsToFile(); err != nil {
		http.Error(w, "Failed to save item", http.StatusInternalServerError)
		return
	}

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

	if err := saveItemsToFile(); err != nil {
		http.Error(w, "Failed to save item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	if err := loadItemsFromFile(); err != nil {
		log.Fatalf("Error loading items from file: %v", err)
	}

	http.HandleFunc("/create", createItem)
	http.HandleFunc("/get", getItem)
	http.HandleFunc("/update", updateItem)
	http.HandleFunc("/delete", deleteItem)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
