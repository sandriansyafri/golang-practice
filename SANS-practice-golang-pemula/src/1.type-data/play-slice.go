package main

import "fmt"

func main() {
	// TODO 1
	// stringArr := [...]string{"a", "b", "c", "d", "e", "f", "g"} // length 7
	// slice1 := stringArr[3:]                                     // d e f g
	// slice2 := stringArr[1:]                                     // b c  e f g
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// TODO 2
	numberArr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceNumber1 := numberArr[3:]
	sliceNumber2 := numberArr[:9]
	sliceNumber3 := numberArr[3:9]
	fmt.Println(sliceNumber1)
	fmt.Println(sliceNumber2)
	fmt.Println(sliceNumber3)
}
