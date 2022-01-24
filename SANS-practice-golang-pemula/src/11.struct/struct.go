package main

import "fmt"

type Person struct {
	Name, Address, Phone string
	Age                  int
}

func main() {
	// TODO 1
	// var person1 Person
	// person1.Name = "Sandrian"
	// person1.Address = "A"
	// person1.Phone = "0812"
	// person1.Age = 32

	// fmt.Println(person1.Name)

	// TODO SECOND STRUCT LITERAL
	// person2 := Person{
	// 	Name:    "Sandrian",
	// 	Address: "AAA",
	// 	Phone:   "09090",
	// 	Age:     23,
	// }
	// fmt.Println(person2)

	// TODO THIRD | Not Recomment
	person3 := Person{"Name", "Address", "Phone", 23}
	fmt.Println(person3)

}
