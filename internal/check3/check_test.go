package check3

import "testing"

func TestCheck3(t *testing.T) {
	if Check3(-1) < 0 {
		t.Fatal()
	}
}
