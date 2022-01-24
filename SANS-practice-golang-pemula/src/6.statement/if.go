package main

import "fmt"

func main() {

	//TODO FULL IF STATMENT
	// name := "Fikri"
	// if name == "Rian" {
	// 	fmt.Println("Hello Rian")
	// } else if name == "Hafid" {
	// 	fmt.Println("Hello Hafid")
	// } else {
	// 	fmt.Println("Hello World!")
	// }

	//TODO SHORT IF STATMENT
	name := "Rian"
	if lengthName := len(name); lengthName > 5 {
		fmt.Println("ok")
	} else {
		fmt.Println("!ok")
	}

}
