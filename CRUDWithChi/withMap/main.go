package main

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Add  string `json:"add"`
}

var Users = map[string]User{}
