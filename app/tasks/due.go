package tasks

import "time"

// Overdue reports whether a task with this due date is late at now. Done
// tasks are never late.
func Overdue(t Task, due, now time.Time) bool {
	return !t.Done && !due.IsZero() && now.After(due)
}
