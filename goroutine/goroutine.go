package main

import (
	"fmt"
	"time"
)

func printSlow() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Slow:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printSlow()
	fmt.Println("Fast")
	time.Sleep(time.Second * 3)
}