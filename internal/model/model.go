package model

import "time"

type Record struct {
	ID, Title, Holder, Region string
	Year                      int
	Amount                    int64
	Status                    string
	Version                   int
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}
type AuditEvent struct {
	ID, RecordID, Actor, Action, Note string
	At                                time.Time
}
type Workflow struct {
	ID, RecordID, Name, State string
	Steps                     []string
	Position                  int
}
type Attachment struct {
	ID, RecordID, Name, ContentType string
	Size                            int64
	Checksum                        string
}
type SearchQuery struct {
	Text, Region, Status string
	Year                 int
}
type Review struct{ RecordID, Reviewer, Decision, Note string }
type ChangeSet struct {
	RecordID, Title, Holder, Region string
	Amount                          int64
	Year                            int
}
