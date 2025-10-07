package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"Suraj", 23},
		{"Amit", 25},
	}

	for _, u := range users {
		fmt.Println(u.Name, u.Age)
	}
}
