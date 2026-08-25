package check5

import "testing"

func TestCheck5(t *testing.T) {
	if Check5(-1) < 0 {
		t.Fatal()
	}
}
