package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[string]User{

	"1":  {"1", "SURAJ", 23},
	"2":  {"2", "MISHRA", 24},
	"3":  {"3", "AMIT", 25},
	"4":  {"4", "SINGH", 26},
	"5":  {"5", "RAHUL", 27},
	"6":  {"6", "PRIYA", 22},
	"7":  {"7", "NEHA", 20},
	"8":  {"8", "ANKIT", 29},
	"9":  {"9", "TANYA", 21},
	"10": {"10", "RAVI", 28},
}

// MAIN
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("API is Working!"))
	}).Methods("GET")

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GET
func getUsers(w http.ResponseWriter, r *http.Request) {

	var allUsers []User // Slice to hold all Users

	for _, user := range users {
		allUsers = append(allUsers, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&allUsers)
}

// POST
func createUser(w http.ResponseWriter, r *http.Request) {

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser) // Decode the JSON request body into the newUser variable

	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	if _, exists := users[newUser.ID]; exists {
		http.Error(w, "User With this ID Already Exists", http.StatusConflict)
		return
	}

	users[newUser.ID] = newUser // Add the new user to the map
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newUser)
}

// PUT
func updateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // Define the URL parameter like "/users/{id}"
	id := params["id"]
	_, exists := users[id]

	if !exists {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	var updateUser User // UpdateUser variable to hold the incoming user
	err := json.NewDecoder(r.Body).Decode(&updateUser)

	if err != nil {
		http.Error(w, "ID is Invalid", http.StatusBadRequest)
		return
	}

	updateUser.ID = id
	users[id] = updateUser // Update the user in the map

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&updateUser)

}

// DELETE
func deleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	if _, exists := users[id]; !exists {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	delete(users, id)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"Message": "User ID Deleted Successfully!"})
}
