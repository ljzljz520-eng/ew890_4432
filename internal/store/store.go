package store

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"heritage/internal/model"
	"time"
)

var buckets = [][]byte{[]byte("records"), []byte("audits"), []byte("workflows"), []byte("attachments")}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error { return s.db.Close() }
func put(tx *bbolt.Tx, b []byte, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), data)
}
func get(tx *bbolt.Tx, b []byte, key string, v any) error {
	d := tx.Bucket(b).Get([]byte(key))
	if d == nil {
		return errors.New("not found")
	}
	return json.Unmarshal(d, v)
}
func (s *Store) SaveRecord(r model.Record) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("records"), r.ID, r) })
}
func (s *Store) LoadRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, []byte("records"), id, &r) })
	return r, e
}
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveAudit(a model.AuditEvent) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("audits"), a.ID, a) })
}
func (s *Store) SaveWorkflow(w model.Workflow) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("workflows"), w.ID, w) })
}
func (s *Store) SaveAttachment(a model.Attachment) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, []byte("attachments"), a.ID, a) })
}
func NewRecord(id, title, holder, region string, year int, amount int64) model.Record {
	now := time.Unix(0, 0).UTC()
	return model.Record{ID: id, Title: title, Holder: holder, Region: region, Year: year, Amount: amount, Status: "draft", Version: 1, CreatedAt: now, UpdatedAt: now}
}
