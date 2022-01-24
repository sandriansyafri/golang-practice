package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name, Expect, Request string
}

func TestSubTestTable(t *testing.T) {
	tests := []struct {
		name    string
		expect  string
		request string
	}{
		{name: "Test1", expect: "Hello Rian", request: "Rian"},
		{name: "Test2", expect: "Hello Hafid", request: "hafid"},
		{name: "Test3", expect: "Hello Fikri", request: "Fikri"},
		{name: "Test4", expect: "Hello Yogi", request: "Yogi"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			assert.Equal(t, test.expect, result)
		})
	}

}

// func TestSubTest(t *testing.T) {
// 	t.Run("Rian", func(t *testing.T) {
// 		result := SayHello("rian")
// 		assert.Equal(t, "Hello Rian", result, "result must be 'Hello Rian'")
// 	})
// 	t.Run("Hafid", func(t *testing.T) {
// 		result := SayHello("Hafid")
// 		assert.Equal(t, "Hello Hafid", result, "result must be 'Hello Hafid'")
// 	})
// }
