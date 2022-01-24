package main

import "fmt"

func main() {

	//1 && (kedua nilai harus memiliki kondisi true -> true , else false )
	//2  || (salah nilai haru memiliki kondisi true -> true, else false)
	//3  ! ( inverse jika kondisi awal bernilai true maka akan diset false )

	// nUTS := 100
	// nUAS := 50
	// nTotal := (nUTS + nUAS) / 2
	// nAbsen := 82

	lulusNilai := true
	lulusAbsensi := true

	lulusSekolah := lulusAbsensi != lulusNilai
	fmt.Println(lulusSekolah)

}
