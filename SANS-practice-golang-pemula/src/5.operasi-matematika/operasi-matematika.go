package main

import "fmt"

func main() {

	// TODO 1
	var x = 10
	var y = 20
	var resultPlus = x + y
	var resultMinus = x - y
	var resultDivide = x / y
	var resultMod = x % 2

	fmt.Println(resultPlus)
	fmt.Println(resultMinus)
	fmt.Println(resultDivide)
	fmt.Println(resultMod)

	// TODO 2
	var i = 10
	// i = i + 20
	i += 20
	fmt.Println(i)

	// TODO 3
	var j = 1
	var k = 2
	j++
	k--

	fmt.Println(j)
	fmt.Println(k)

	// TODO 4
	var negativeNumber = -200
	fmt.Println(negativeNumber)

}
