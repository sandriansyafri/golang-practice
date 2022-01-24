package main

import "fmt"

func main() {

	/**
		1. melakukan statment perulangan
		2. mulai dari  1
		3. tampilkan  1 - 10
		4. selesai
	**/

	// TODO 1
	// i := 1 // init
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// TODO 2
	// name := "SANDRIAN"
	// lengthName := len(name)

	// for i := 0; i < lengthName; i++ {
	// 	fmt.Println(string(name[i]))
	// }

	// TODO 3
	// simpleArr := [...]int{1, 2, 3, 4, 5, 6}
	// for i, val := range simpleArr {
	// 	fmt.Println(i, val)
	// }

	// simpleSlice := []int{1, 2, 3, 4, 5, 6}
	// for _, value := range simpleSlice {
	// 	fmt.Println(value)
	// }

	// TODO 4
	simpleMap := make(map[string]string)
	simpleMap["name"] = "a"
	simpleMap["username"] = "a"
	simpleMap["email"] = "a"
	simpleMap["passsword"] = "a"

	for key, value := range simpleMap {
		fmt.Println(key, value)
	}

}
