package main

import "fmt"

func findTheMinimum(num []int) int {
	min := num[0]
	for _, value := range num {
		if min > value {
			min = value
		}

	}
	return min
}

func main() {
	n := []int{5, 8, 2, 8, 0, -3}
	fmt.Println(findTheMinimum(n))
}
