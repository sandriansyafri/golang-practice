package main

import "fmt"

type HasName interface {
	getName() string
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

func printResult(print HasName) {
	fmt.Println("This is", print.getName())
}

func main() {
	cat := Animal{Name: "CAT"}

	printResult(cat)

}
