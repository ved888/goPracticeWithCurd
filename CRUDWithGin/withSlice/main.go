package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var Users []User

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	Users = append(Users, user)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func getAll(c *gin.Context) {
	c.Header("contant-Type", "application/json")

	c.JSON(http.StatusOK, gin.H{"users": Users})
}

func getById(c *gin.Context) {
	id := c.Param("id")

	for _, user := range Users {
		if user.Id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusNotFound, gin.H{"error": "uder not found"})
}

func updateUser(c *gin.Context) {
	id := c.Query("id")
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	index := -1

	for i, u := range Users {
		if u.Id == id {
			index = i
		}
	}
	Users[index].Name = user.Name
	Users[index].Add = user.Add

	c.JSON(http.StatusOK, gin.H{"message": "updated user"})

}

func delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	newUser := []User{}

	for _, user := range Users {
		if user.Id != id {
			newUser = append(newUser, user)
		}
	}
	Users = newUser
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "deleted user"})
}

func main() {
	r := gin.Default()

	r.POST("/create", createUser)
	r.GET("/getAll", getAll)
	r.GET("/get/:id", getById)
	r.PUT("update", updateUser)
	r.DELETE("/delete/:id", delete)

	fmt.Println("server start on port 8080")
	log.Fatal(r.Run(":8080"))
}
