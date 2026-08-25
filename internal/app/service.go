package app

import (
	"errors"
	"fmt"
	"heritage/internal/model"
	"heritage/internal/store"
)

type Service struct {
	st *store.Store
}

func New(st *store.Store) *Service { return &Service{st: st} }
func (s *Service) Register(id, title, holder, region string, year int, amount int64) (model.Record, error) {
	if id == "" || title == "" {
		return model.Record{}, errors.New("id and title required")
	}
	r := store.NewRecord(id, title, holder, region, year, amount)
	if e := s.st.SaveRecord(r); e != nil {
		return r, e
	}
	s.st.SaveAudit(model.AuditEvent{ID: id + "-register", RecordID: id, Action: "register"})
	return r, nil
}
func (s *Service) Review(id, actor, decision, note string) error {
	r, e := s.st.LoadRecord(id)
	if e != nil {
		return e
	}
	if decision != "approve" && decision != "reject" {
		return errors.New("invalid decision")
	}
	if decision == "approve" {
		r.Status = "approved"
	} else {
		r.Status = "rejected"
	}
	r.Version++
	if e = s.st.SaveRecord(r); e != nil {
		return e
	}
	return s.st.SaveAudit(model.AuditEvent{ID: fmt.Sprintf("%s-review-%d", id, r.Version), RecordID: id, Actor: actor, Action: decision, Note: note})
}
func (s *Service) Change(c model.ChangeSet) error {
	r, e := s.st.LoadRecord(c.RecordID)
	if e != nil {
		return e
	}
	if c.Title != "" {
		r.Title = c.Title
	}
	if c.Holder != "" {
		r.Holder = c.Holder
	}
	if c.Region != "" {
		r.Region = c.Region
	}
	if c.Year > 0 {
		r.Year = c.Year
	}
	if c.Amount >= 0 {
		r.Amount = c.Amount
	}
	r.Version++
	return s.st.SaveRecord(r)
}
func (s *Service) Archive(id string) error {
	r, e := s.st.LoadRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "approved" {
		return errors.New("record must be approved")
	}
	r.Status = "archived"
	r.Version++
	return s.st.SaveRecord(r)
}
func (s *Service) Summary(id string) string {
	r, e := s.st.LoadRecord(id)
	if e != nil {
		return e.Error()
	}
	return fmt.Sprintf("%s|%s|%s|%d", r.ID, r.Title, r.Status, r.Amount)
}
func (s *Service) Search(q model.SearchQuery) ([]model.Record, error) {
	all, e := s.st.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range all {
		if q.Region != "" && r.Region != q.Region {
			continue
		}
		if q.Status != "" && r.Status != q.Status {
			continue
		}
		if q.Year > 0 && r.Year != q.Year {
			continue
		}
		if q.Text != "" && r.Title != q.Text && r.Holder != q.Text {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}
func (s *Service) Import(rows []model.Record) error {
	for _, r := range rows {
		if r.ID == "" {
			continue
		}
		if e := s.st.SaveRecord(r); e != nil {
			return e
		}
	}
	return nil
}
