package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		result := SayHello("tian")
		assert.Equal(t, "Hello Rian", result, "-")
	})
}
