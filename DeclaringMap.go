package main

import "fmt"

func declaringMap() {
	x := make(map[string]int)
	x["key"] = 10
	x["Josh"] = 24
	x["Kuse"] = 30

	name, ok := x["key"]
	fmt.Println(name, ok)
}

func main() {
	declaringMap()
	fmt.Println("hello")

}
