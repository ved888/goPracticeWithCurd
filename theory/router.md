# Go Routers with Gin, Mux, and Chi

## Gin Router

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create user
	r.POST("/users", createUser)

	// Get all users
	r.GET("/users", getUsers)

	// Get user by ID
	r.GET("/users/:id", getUserByID)

	// Update user by ID
	r.PUT("/users/:id", updateUser)

	// Delete user by ID
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}
```

## Mux Router

```go
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Create user
	r.HandleFunc("/users", createUser).Methods("POST")

	// Get all users
	r.HandleFunc("/users", getUsers).Methods("GET")

	// Get user by ID
	r.HandleFunc("/users/{id}", getUserByID).Methods("GET")

	// Update user by ID
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	// Delete user by ID
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
```

## Chi Router

```go
package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// Create user
	r.Post("/users", createUser)

	// Get all users
	r.Get("/users", getUsers)

	// Get user by ID
	r.Get("/users/{id}", getUserByID)

	// Update user by ID
	r.Put("/users/{id}", updateUser)

	// Delete user by ID
	r.Delete("/users/{id}", deleteUser)

	http.ListenAndServe(":8080", r)
}
```



# Go Routers with Gin, Mux, and Chi

| **Gin Router**  | **Mux Router**  | **Chi Router**  |
|-----------------|-----------------|-----------------|
| ```go           | ```go           | ```go           |
| `r := gin.Default()` | `r := mux.NewRouter()` | `r := chi.NewRouter()` |
| `r.POST("/users", createUser)` | `r.HandleFunc("/users", createUser).Methods("POST")` | `r.Post("/users", createUser)` |
| `r.GET("/users", getUsers)` | `r.HandleFunc("/users", getUsers).Methods("GET")` | `r.Get("/users", getUsers)` |
| `r.GET("/users/:id", getUserByID)` | `r.HandleFunc("/users/{id}", getUserByID).Methods("GET")` | `r.Get("/users/{id}", getUserByID)` |
| `r.PUT("/users/:id", updateUser)` | `r.HandleFunc("/users/{id}", updateUser).Methods("PUT")` | `r.Put("/users/{id}", updateUser)` |
| `r.DELETE("/users/:id", deleteUser)` | `r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")` | `r.Delete("/users/{id}", deleteUser)` |
| `r.Run(":8080")` | `http.ListenAndServe(":8080", r)` | `http.ListenAndServe(":8080", r)` |
| ```             | ```             | ```             |


# Go Routers with Gin, Mux, and Chi

| **Gin Router**                                    | **Mux Router**                                    | **Chi Router**                                    |
|---------------------------------------------------|---------------------------------------------------|---------------------------------------------------|
| ```gin                                             | ```mux                                             | ```chi                                             |
| r := gin.Default()                                | r := mux.NewRouter()                             | r := chi.NewRouter()                             |
| r.POST("/users", createUser)                      | r.HandleFunc("/users", createUser).Methods("POST")| r.Post("/users", createUser)                     |
| r.GET("/users", getUsers)                         | r.HandleFunc("/users", getUsers).Methods("GET")   | r.Get("/users", getUsers)                        |
| r.GET("/users/:id", getUserByID)                  | r.HandleFunc("/users/{id}", getUserByID).Methods("GET") | r.Get("/users/{id}", getUserByID)             |
| r.PUT("/users/:id", updateUser)                   | r.HandleFunc("/users/{id}", updateUser).Methods("PUT")   | r.Put("/users/{id}", updateUser)                |
| r.DELETE("/users/:id", deleteUser)                | r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")| r.Delete("/users/{id}", deleteUser)            |
| r.Run(":8080")                                    | http.ListenAndServe(":8080", r)                   | http.ListenAndServe(":8080", r)                  |
| ```                                               | ```                                               | ```                                               |
