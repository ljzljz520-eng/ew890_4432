package app

import (
	"heritage/internal/model"
	"heritage/internal/store"
	"testing"
)

func testService(t *testing.T) *Service {
	t.Helper()
	s, e := store.Open(t.TempDir() + "/x.db")
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return New(s)
}
func TestWorkflowCreateReviewArchive(t *testing.T) {
	s := testService(t)
	if _, e := s.Register("a", "A", "H", "R", 1900, 10); e != nil {
		t.Fatal(e)
	}
	if e := s.Review("a", "u", "approve", ""); e != nil {
		t.Fatal(e)
	}
	if e := s.Archive("a"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowSearchUpdatePublish(t *testing.T) {
	s := testService(t)
	s.Register("a", "A", "H", "R", 1900, 10)
	r, e := s.Search(model.SearchQuery{Text: "A"})
	if e != nil || len(r) != 1 {
		t.Fatal(e, len(r))
	}
	if e = s.Change(model.ChangeSet{RecordID: "a", Amount: 20}); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowImportReport(t *testing.T) {
	s := testService(t)
	if e := s.Import([]model.Record{{ID: "a", Title: "A"}, {Title: "skip"}}); e != nil {
		t.Fatal(e)
	}
	r, e := s.Search(model.SearchQuery{})
	if e != nil || len(r) != 1 {
		t.Fatal(e, len(r))
	}
}
func Test890BusinessRegression(t *testing.T) {
	s := testService(t)
	s.Register("one", "One", "H", "R", 1, 11)
	s.Register("two", "Two", "H", "R", 2, 22)
	r, e := s.st.LoadRecord("two")
	if e != nil {
		t.Fatal(e)
	}
	if r.Amount != 22 {
		t.Fatalf("amount=%d", r.Amount)
	}
}
