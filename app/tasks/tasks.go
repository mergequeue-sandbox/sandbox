// Package tasks is a small stand-in for real code. It's here so the sandbox
// can show a semantic conflict: two PRs that each pass CI, but break
// develop once both are merged.
package tasks

import "fmt"

// Task is a to-do item.
type Task struct {
	ID   int
	Name string
	Done bool
}

// Title is how a task is shown in lists and notifications.
func Title(t Task) string {
	return fmt.Sprintf("#%d %s", t.ID, t.Name)
}
