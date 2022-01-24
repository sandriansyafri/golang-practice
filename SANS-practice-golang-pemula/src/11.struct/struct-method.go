package main

import "fmt"

type Person struct {
	Name, Address, Phone string
	Age                  int
}

//TODO MANUAL METHOD
// func  hi(person Person, name string) {
// 	fmt.Println("Hello", name, "MyName is", person.Name)
// }

//TODO STRUCT METHOD
func (person Person) hi(name string) {
	fmt.Println("Hello ", name, "I am", person.Name)
}

func main() {

	person1 := Person{
		Name:    "Sandrian",
		Address: "Jambi",
		Phone:   "0800xxx",
		Age:     23,
	}

	person1.hi("Boy")

}
