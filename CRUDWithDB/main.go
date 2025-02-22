package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func dbConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Host := os.Getenv("HOST")
	DbPort := os.Getenv("DBPORT")
	DbName := os.Getenv("DBNAME")
	DbUser := os.Getenv("DBUSER")
	Password := os.Getenv("PASSWORD")

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", Host, DbPort, DbUser, DbName, Password)
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		log.Fatal("db connection is failed")
	}

	err = db.Ping()
	if err != nil {
		log.Println("problem in db connection", err)
		panic(err)
	}
	DB = db
	log.Println("db connection is successfull")
}

type User struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	_, err = DB.Exec(`
	            INSERT INTO users(
				       first_name,
					   last_name,
					   email
				)
			    VALUES(
				       $1,
					   $2,
					   $3
				);`,
		user.FirstName,
		user.LastName,
		user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	var user []User
	err := DB.Select(&user, `Select id,first_name,last_name,email from users where deleted_at is null`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user User
	err := DB.Get(&user, `select id,first_name,last_name,email from users where id=$1 and deleted_at is null`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No record found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = DB.Exec(
		`update users
	                set 
					   first_name=$1,
					   last_name=$2,
					   email=$3,
					   updated_at=Now()
				 where id=$4`,
		user.FirstName,
		user.LastName,
		user.Email,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

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
	dbConnection()
	// make router
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/getAll", getAllUser)
	http.HandleFunc("/getById", getUserById)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)
	port := os.Getenv("PORT")

	fmt.Printf("server start on port:%s", port)
	http.ListenAndServe(":"+port, nil)
}
