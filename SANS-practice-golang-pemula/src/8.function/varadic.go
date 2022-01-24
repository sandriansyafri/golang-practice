package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total

}

func main() {
	numbers := []int{1, 2, 3, 5, 10}
	result := sumAll(numbers...)
	fmt.Println(result)
}
