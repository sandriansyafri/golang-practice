package sans_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {

	//windows == windows = true
	//windows != windows = false
	if runtime.GOOS == "darwins" {
		t.Skip()
	}

	fmt.Println("....")
}
