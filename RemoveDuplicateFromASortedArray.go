package main

import "fmt"

func removeDuplicate(num []int) []int {
	count := 1
	for i := 0; i < len(num); i++ {
		if num[i] == num[count] {
			num = append(num[:i], num[count:]...)
		}
		count++

	}
	return num

}

func main() {
	num := []int{1, 2, 2, 4, 5, 5, 7, 8, 8}
	fmt.Println(removeDuplicate(num))

}
