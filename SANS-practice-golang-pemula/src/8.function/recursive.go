package main

import "fmt"

func factorialLoop(value int) int {
	// TODO INCREMENT
	// result := 1
	// for i := 1; i <= value; i++ {
	// 	result *= i
	// }
	// return result

	//TODO DECREMENT
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialLoopRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoopRecursive(value-1)
	}
}

func main() {
	// loopFactorialFor := factorialLoop(5)
	// fmt.Println(loopFactorialFor)
	loopFactorialRecursive := factorialLoopRecursive(3)
	fmt.Println(loopFactorialRecursive)
}
