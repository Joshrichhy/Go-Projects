package main

import "fmt"

func sumDigitToMeetTarget(num []int, target int) (firstIndex int, secondIndex int) {

	for i := 0; i < len(num); i++ {
		for j := 0; i < len(num); j++ {
			if num[i]+num[j] == target {
				return i, j
			}
		}
	}
	return 0, 0
}

func main() {
	numArray := []int{1, 2, 4, 5, 6, 3, 7}
	fmt.Println(sumDigitToMeetTarget(numArray, 8))

}
