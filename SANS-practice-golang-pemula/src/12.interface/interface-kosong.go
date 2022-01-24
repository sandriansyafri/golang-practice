package main

import "fmt"

type Any interface{}

func test(i int) Any {
	if i == 1 {
		return i
	} else if i == 2 {
		return true
	} else {
		return "string"
	}
}

func main() {
	x := test(3)
	fmt.Println(x)
}
