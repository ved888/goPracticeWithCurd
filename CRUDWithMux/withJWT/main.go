package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

// Secret key for JWT signing
var secretKey = []byte("secret-key")

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]User{
	"Chek": {"Chek", "12345"},
}

func CreateTocken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	return token.SignedString(secretKey)

}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")

}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeder := r.Header.Get("Authorization")
		if authHeder == "" || len(authHeder) <= len("Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := authHeder[len("Bearer "):]
		claims, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Printf("Authenticated user: %v\n", claims["username"])
		next.ServeHTTP(w, r)
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	storeUser, exist := users[user.UserName]
	if !exist || storeUser.Password != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := CreateTocken(user.UserName)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	users[user.UserName] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created"})
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userList []User

	for _, user := range users {
		userList = append(userList, User{UserName: user.UserName, Password: user.Password})
	}
	json.NewEncoder(w).Encode(userList)
}

func getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	username := param["username"]
	user, exist := users[username]
	if !exist {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(User{UserName: user.UserName, Password: user.Password})
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := r.URL.Query().Get("username")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	_, exist := users[username]
	if !exist {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	users[username] = user

	json.NewEncoder(w).Encode(map[string]string{"message": "user updated"})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]

	_, exist := users[username]
	if !exist {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(users, username)
	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted"})
}

func main() {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/create", createUser).Methods("POST")

	r.HandleFunc("/getAll", AuthMiddleware(getAll)).Methods("GET")
	r.HandleFunc("/getByID/{username}", AuthMiddleware(getById)).Methods("GET")
	r.HandleFunc("/update", AuthMiddleware(updateUser)).Methods("PUT")
	r.HandleFunc("/delete/{username}", AuthMiddleware(deleteUser)).Methods("DELETE")

	fmt.Println("Starting the server on port 4000...")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
