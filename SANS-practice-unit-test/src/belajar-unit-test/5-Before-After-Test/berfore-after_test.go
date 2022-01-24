package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE")
	m.Run()
	fmt.Println("AFTER")
}

func TestSayHelloRian(t *testing.T) {
	result := SayHello("Rian")
	assert.Equal(t, "Hello Rian", result, "must be 'Hello Rian'")
	fmt.Println("DONE")
}
