package main

import (
	"17-pkg-init/db"
	_ "17-pkg-init/db"
	"fmt"
)

func main() {
	fmt.Println(db.GetConnection())
}
