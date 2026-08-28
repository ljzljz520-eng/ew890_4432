package check4

import "testing"

func TestCheck4(t *testing.T) {
	if Check4(-1) < 0 {
		t.Fatal()
	}
}
