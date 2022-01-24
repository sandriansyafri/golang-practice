package main

import "fmt"

func sayHi(firstName string, lastName string) string {
	return "Hello " + firstName + " " + lastName
}

// TODO FIRST
func main() {
	fmt.Println(sayHi("Sandrian", "Hafid"))
}
