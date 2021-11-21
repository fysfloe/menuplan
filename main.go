package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Ingredient struct
type Ingredient struct {
	gorm.Model
	Name string
}

// User struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

// Recipe struct
type Recipe struct {
	gorm.Model
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

var Users []User

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint hit: homePage")
}

func userList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: userList")
	json.NewEncoder(w).Encode(Users)
}

func handleRequests() {
	http.HandleFunc("/users", userList)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Users = []User{
		User{Name: "Florian Hofinger", Email: "hofinger.flo@gmail.com"},
	}

	handleRequests()
}
