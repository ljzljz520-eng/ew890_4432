package check6

import "testing"

func TestCheck6(t *testing.T) {
	if Check6(-1) < 0 {
		t.Fatal()
	}
}
