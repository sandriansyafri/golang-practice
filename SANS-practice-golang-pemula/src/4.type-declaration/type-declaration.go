package main

import "fmt"

type KTP string
type isMale bool

func main() {

	var noKTP KTP = "15100"
	var gender isMale = true
	fmt.Println(noKTP)
	fmt.Println(gender)

}
