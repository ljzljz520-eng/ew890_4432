package model

import "testing"

func TestRecordDefaults(t *testing.T) {
	r := Record{Status: "draft"}
	if r.Status != "draft" {
		t.Fatal()
	}
}
