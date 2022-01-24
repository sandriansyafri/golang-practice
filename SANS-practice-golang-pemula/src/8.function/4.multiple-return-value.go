package main

import "fmt"

func sayHi() (string, string) {
	return "Sandrian", "Syafri"
}

func main() {
	firstName, lastName := sayHi()
	fmt.Println(firstName, lastName)
}
