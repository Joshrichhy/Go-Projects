package main

import "fmt"

func average(num []int) float64 {
	total := 0
	for _, value := range num {
		total += value

	}
	return float64(total / len(num))
}

func smallestValue(num []int) int {
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
	fmt.Println(smallestValue(n))

}
