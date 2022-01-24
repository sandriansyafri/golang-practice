package main

import "fmt"

func main() {

	count := 0
	name := "Rian"
	increment := func() {
		name = "Ok"
		count++
	}
	fmt.Println(name)
	increment()
	fmt.Println(name)

}
