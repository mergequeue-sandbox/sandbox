// Package projects groups tasks.
package projects

// Archived reports whether a project with this status is hidden from lists.
func Archived(status string) bool { return status == "archived" }
