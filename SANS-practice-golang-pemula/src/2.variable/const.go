package main

import "fmt"

func main() {
	const fullName = "Sandrian Syafri"
	const age = 23
	fmt.Println(fullName)
	fmt.Println(age)

	const (
		firstName = "Sandrian"
		lastName  = "Syafri"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
