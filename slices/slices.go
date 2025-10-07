package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "mango")
	fmt.Println("Fruits:", fruits)
}
