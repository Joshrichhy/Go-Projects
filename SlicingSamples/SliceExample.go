package main

import "fmt"

func addingToASlice() {
	slice1 := []int{2, 3, 4, 5, 6, 7}
	slice2 := append(slice1, 1, 4)
	fmt.Println(slice1, "second slice is ", slice2)
}

func copyingFromAnExistingArrayToAnother() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 4)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func main() {

	//x := make([]float64, 5)
	//y := []int{1, 2, 3, 4, 5}
	//arr := y[1:4]
	//fmt.Println(arr, "former length is ", x)
	//addingToASlice()
	//copyingFromAnExistingArrayToAnother()

}
