package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
}

var DB *sqlx.DB

func DbConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	name := os.Getenv("DBNAME")
	pass := os.Getenv("PASSWORD")

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, pass)
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		log.Fatal("db is not connection")
	}

	err = db.Ping()
	if err != nil {
		log.Println("problem in db connection", err)
		panic(err)
	}
	DB = db
	fmt.Println("db connection is successful")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	_, err = DB.Exec(`insert into users(
                     first_name,
                     last_name,
                     email
                     )values(
                     $1,
                     $2,
                     $3)
                     `,
		user.FirstName,
		user.LastName,
		user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contant-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func getAll(w http.ResponseWriter, r *http.Request) {
	var user []User

	err := DB.Select(&user, `select 
                          first_name,
                          last_name,
                          email,
                          id
                    from users
                          where deleted_at is null`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contant-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var user User

	err := DB.Get(&user, `select
                             first_name,
                             last_name,
                             email
                        from users
                        where id=$1 and deleted_at is null`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No record found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Contant-TYpe", "application/json")
	json.NewEncoder(w).Encode(user)
}

func updateUSer(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	_, err = DB.Exec(`update users
                 set
                    first_name=$1,
                    last_name=$2,
                    email=$3,
                    updated_at=NOW()
               where id=$4
                    `,
		user.FirstName,
		user.LastName,
		user.Email,
		id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contant-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func delewteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	_, err := DB.Exec(`update users
              set
                 deleted_at=now()
              where id=$1   
                 `, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func main() {
	DbConnection()

	// mux router
	r := mux.NewRouter()
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/getll", getAll).Methods("GET")
	r.HandleFunc("/getById", getByID).Methods("GET")
	r.HandleFunc("/update/{id}", updateUSer).Methods("PUT")
	r.HandleFunc("/delete/{id}", delewteUser).Methods("DELETE")
	port := os.Getenv("PORT")
	fmt.Printf("server start on port:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
