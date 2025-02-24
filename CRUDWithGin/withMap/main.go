package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var Users = map[string]User{}

func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erroe": "invalid request body"})
		return
	}
	Users[user.Id] = user
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func getAll(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"users": Users})
}

func getById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plese provide valid user id"})
		return
	}

	user, exist := Users[id]
	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func getID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	for _, value := range Users {
		if value.Id == id {
			c.JSON(http.StatusOK, gin.H{"user": value})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func update(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	_, found := Users[id]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user is not found"})
		return
	}
	Users[id] = User{
		Id:   user.Id,
		Name: user.Name,
		Add:  user.Add,
	}

	c.Status(http.StatusNoContent)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	_, found := Users[id]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	delete(Users, id)
	c.Status(http.StatusNoContent)

}

func main() {
	r := gin.Default()

	r.POST("/create", createUser)
	r.GET("/getAll", getAll)
	r.GET("/get", getById)
	r.GET("/getId/:id", getID)
	r.PUT("/update/:id", update)
	r.DELETE("/delete/:id", deleteUser)

	r.Run(":8080")
}
