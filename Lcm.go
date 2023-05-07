package main

import "fmt"

func main() {

	var number = 8
	var count = 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count += 1
		}
	}

	const count2 = 4
	var lcm [count2]int8
	var counter int = 0

	for i := 2; i <= number; i++ {
		if number%i == 0 {
			number = number / i
			lcm[counter] = int8(i)
			i = 1
			counter++
		}

	}

	fmt.Println(lcm)

}
