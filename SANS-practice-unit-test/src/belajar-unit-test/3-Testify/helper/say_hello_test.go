package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSayHelloRian(t *testing.T) {
	result := SayHello("Rian")
	// assert.Equal(t, "Hello rian", result, "Error, result must be 'Hello Rian'")
	require.Equal(t, "Hello rian", result, "Error, result must be 'Hello Rian'")
	fmt.Println("DONE")
}
