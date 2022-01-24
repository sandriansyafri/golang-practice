package helper

import (
	"fmt"
	"testing"
)

func TestSayHelloRian(t *testing.T) {
	result := SayHello("Rian")

	//! FAIL
	if result != "Hello Rian" {
		panic("ERROR!")
	}
	//* SUCCESS
	fmt.Println("DONE")
}
func TestSayHelloHafid(t *testing.T) {
	result := SayHello("hafid")

	if result != "Hello Hafid" {
		panic("ERROR!")
	}

	fmt.Println("DONE")
}
