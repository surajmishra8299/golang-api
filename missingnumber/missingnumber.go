package main

import "fmt"

func isMissing(nums []int) int {

	n := len(nums) + 1
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}

func main() {

	res := isMissing([]int{1, 2, 3, 6, 4, 7, 9, 5})
	fmt.Println("Missing Number:", res)
}
