package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  uint
}

func (p Person) getPersonInfo() string {
	return p.Name + " " + strconv.Itoa(int(p.Age))

}

func main() {
	person := new(Person)
	person1 := new(Person)
	person.Name = "Joshua"
	person.Age = 20
	person1.Name = "Oluwakuse"
	person1.Age = 12
	fmt.Println(person.getPersonInfo())
	fmt.Println(person1.getPersonInfo())
}
