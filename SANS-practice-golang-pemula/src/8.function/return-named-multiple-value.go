package main

import "fmt"

func sayHi() (firstName string, lastName string) {
	firstName = "sandrian"
	lastName = "syafri"
	return
}

func main() {
	a, b := sayHi()
	fmt.Println(a, b)
}
