package main

import "fmt"

func countDigit(words []string) int {
	var count = 0
	for _, word := range words {

		for i := 0; i < len(word); i++ {
			if word[i] >= 48 && word[i] <= 57 {

				count += 1
				fmt.Println(count)
			}

		}

	}
	return count
}

func main() {
	words := []string{"whfb432", "12dfdg", "321asfv"}
	fmt.Println(countDigit(words))

}
