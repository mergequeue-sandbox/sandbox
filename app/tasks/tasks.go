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

// DisplayTitle is how a task is shown in lists and notifications. Done
// tasks get a check mark.
func DisplayTitle(t Task) string {
	if t.Done {
		return fmt.Sprintf("✓ #%d %s", t.ID, t.Name)
	}
	return fmt.Sprintf("#%d %s", t.ID, t.Name)
}
