package main

import "fmt"

func main() {

	var (
		num32   int32 = 100000
		c_num64 int64 = int64(num32)

		name    = "Rian"[2]
		eString = string(name)
	)

	fmt.Println(num32)
	fmt.Println(c_num64)

	fmt.Println(eString)
}
