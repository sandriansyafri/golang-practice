package helper

import (
	"fmt"
	"testing"
)

func TestSayHelloRian(t *testing.T) {
	result := SayHello("Rian")

	if result != "Hello Rian" {
		t.Error()
	}

	fmt.Println("DONE")
}
func TestSayHelloHafid(t *testing.T) {
	result := SayHello("Hafid!")

	if result != "Hello Hafid" {
		// t.Fail()
		// t.FailNow()
		t.Error("ERROR, test must be 'Hello Hafid'")
		// t.Fatal("ERROR, test must be 'Hello Hafid'")
	}

	fmt.Println(result)
}
