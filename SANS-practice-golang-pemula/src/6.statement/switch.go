package main

import "fmt"

func main() {

	// TODO FIRST
	// name := "Yogi"
	// switch name {
	// case "Rian":
	// 	fmt.Println("Hello Rian")
	// 	fmt.Println("Hello Rian")
	// case "Hafid":
	// 	fmt.Println("Hello Hafid")
	// 	fmt.Println("Hello Hafid")
	// case "Fikri":
	// 	fmt.Println("Hello Fikri")
	// 	fmt.Println("Hello Fikri")
	// default:
	// 	fmt.Println("Hello World!")
	// 	fmt.Println("Hello World!")
	// }

	// TODO SECOND
	// switch name := "sandrian"; len(name) > 5 {
	// case true:
	// 	fmt.Println("Sandrian")
	// case false:
	// 	fmt.Println("Hello Wolrd!")
	// }
	// TODO THIRD
	name := "sdafjklsafjd"
	switch {
	case name == "sandrian":
		fmt.Println("Hello sandrian")
	case name == "hafid":
		fmt.Println("Hello hafid")
	case name == "fikri":
		fmt.Println("Hello fikri")
	default:
		fmt.Println("Hello WOrl!!")
	}
}
