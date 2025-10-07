package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{"Suraj", 23}
	fmt.Println(u.Name, u.Age)
}
