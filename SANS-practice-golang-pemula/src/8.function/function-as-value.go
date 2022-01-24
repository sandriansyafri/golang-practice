package main

import "fmt"

func sayHi(name string) string {
	return "Hello " + name
}

func main() {
	getSayHi := sayHi
	fmt.Println(getSayHi("Sadrian"))
}
