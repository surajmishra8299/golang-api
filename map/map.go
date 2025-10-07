package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Suraj",
		"city": "Prayagraj",
	}
	fmt.Println("Name:", person["name"])
}
