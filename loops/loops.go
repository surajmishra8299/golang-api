package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}

	j := 1
	for j <= 3 {
		fmt.Println("j =", j)
		j++
	}
}
