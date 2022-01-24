package main

import "fmt"

type Person struct {
	Name, Username, Email string
}

func main() {
	// TODO 1
	// person1 := Person{"A", "A", "A"}
	// person2 := &person1
	// person3 := &person1

	// person2.Name = "Edited"
	// fmt.Println(person3)
	// *person2 = Person{"edited", "edited", "edited"}

	// fmt.Println(person1)
	// fmt.Println(person2)

	// TODO 2
	person3 := new(Person)
	*person3 = Person{"A", "B", "C"}

	fmt.Println(person3)

}
