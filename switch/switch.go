package main

import "fmt"

func main() {
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Another day")
	}
}
