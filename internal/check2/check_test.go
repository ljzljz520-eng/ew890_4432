package check2

import "testing"

func TestCheck2(t *testing.T) {
	if Check2(-1) < 0 {
		t.Fatal()
	}
}
