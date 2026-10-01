package tasks

import (
	"testing"
	"time"
)

func TestOverdue(t *testing.T) {
	now := time.Date(2026, 10, 2, 9, 0, 0, 0, time.UTC)
	if !Overdue(Task{ID: 1}, now.Add(-time.Hour), now) {
		t.Error("an open task past its due date should be overdue")
	}
	if Overdue(Task{ID: 2, Done: true}, now.Add(-time.Hour), now) {
		t.Error("a done task is never overdue")
	}
}
