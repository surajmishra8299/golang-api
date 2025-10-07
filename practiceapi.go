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
	"7":  {"7", "NEHA", 24},
	"8":  {"8", "ANKIT", 29},
	"9":  {"9", "TANYA", 23},
	"10": {"10", "RAVI", 28},
}

// MAIN

func main1() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("API IS PROPERLY WORKING!"))
	}).Methods("GET")

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("Server runing at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GET

func getUsers(w http.ResponseWriter, r *http.Request) {

	var allUsers []User

	for _, user := range users {
		allUsers = append(allUsers, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}

// POST
func createUser(w http.ResponseWriter, r *http.Request) {

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	if _, exists := users[newUser.ID]; exists {
		http.Error(w, "User with this ID already exists", http.StatusConflict)
		return
	}

	users[newUser.ID] = newUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// PUT

func updateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	_, exists := users[id]

	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updateUser User
	err := json.NewDecoder(r.Body).Decode(&updateUser)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	updateUser.ID = id
	users[id] = updateUser

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(updateUser)
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Message deletes successfully"})

}
