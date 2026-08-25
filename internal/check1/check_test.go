package check1

import "testing"

func TestCheck1(t *testing.T) {
	if Check1(-1) < 0 {
		t.Fatal()
	}
}
