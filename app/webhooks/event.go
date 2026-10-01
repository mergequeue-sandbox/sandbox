// Package webhooks sends task events to the URLs people register.
package webhooks

// Event is something that happened to a task.
type Event struct {
	Kind   string // "task.created", "task.completed", …
	TaskID int
}
