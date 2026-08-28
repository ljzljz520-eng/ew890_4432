package store

import (
	"heritage/internal/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	s.SaveRecord(NewRecord("x", "X", "H", "R", 1, 7))
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, e := s.LoadRecord("x")
	if e != nil || r.Amount != 7 {
		t.Fatal(e, r)
	}
}
func TestStoreCollections(t *testing.T) {
	s := mustOpen(t)
	defer s.Close()
	if e := s.SaveAudit(model.AuditEvent{ID: "a"}); e != nil {
		t.Fatal(e)
	}
	if e := s.SaveWorkflow(model.Workflow{ID: "w"}); e != nil {
		t.Fatal(e)
	}
	if e := s.SaveAttachment(model.Attachment{ID: "f"}); e != nil {
		t.Fatal(e)
	}
}
func mustOpen(t *testing.T) *Store {
	s, e := Open(t.TempDir() + "/x")
	if e != nil {
		t.Fatal(e)
	}
	return s
}
