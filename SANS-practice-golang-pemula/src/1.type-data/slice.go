package main

import "fmt"

func main() {
	//TODO FIRST
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// strings := [...]string{"satu", "dua", "tiga", "empat"}
	// sliceString1 := strings[1:]
	// fmt.Println(sliceString1)

	sliceNumbers1 := numbers[1:] // 5, 6
	// sliceNumbers2 := numbers[2:] // 1,2,3,4
	// sliceNumbers3 := numbers[:4] // 1,2,3,4
	// sliceNumbers4 := numbers[:]  // 1,2,3,4,5,6

	fmt.Println(sliceNumbers1)
	return

	// fmt.Println(sliceNumbers2)
	// fmt.Println(sliceNumbers3)

	// TODO SECOND
	// months := [...]string{
	// 	"Jan",
	// 	"Feb",
	// 	"Mar",
	// 	"Apr",
	// 	"May",
	// 	"Jul",
	// 	"Jun",
	// 	"Aug",
	// 	"Sep",
	// 	"Okt",
	// 	"Nov",
	// 	"Des",
	// }

	// slice1 := months[4:8]
	// fmt.Println("CONTENT : ", slice1)
	// fmt.Println("LEN :", len(slice1))
	// fmt.Println("CAP: ", cap(slice1))

	// slice2 := months[2:5]
	// fmt.Println(slice2)
	// slice3 := append(slice2, "updated")
	// fmt.Println(slice3)

	// slice3[1] = "ubah"
	// fmt.Println(slice3)

	// fmt.Println(slice2)
	// fmt.Println(months)

	// TODO THIRD
	// newSlice := make([]string, 2)
	// newSlice[0] = "satu"
	// newSlice[1] = "dua"

	// copySlice := make([]string, len(newSlice))
	// fmt.Println(copySlice)
	// copy(copySlice, newSlice)
	// fmt.Println(newSlice)
	// fmt.Println(copySlice)

	// TODO FOUR
	iniArray := [5]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println((iniArray))
	fmt.Println((iniSlice))

}
