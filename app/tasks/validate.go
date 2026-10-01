package tasks

import (
	"errors"
	"strings"
)

// Validate rejects tasks that would show up as a blank line.
func Validate(t Task) error {
	if strings.TrimSpace(t.Name) == "" {
		return errors.New("a task needs a name")
	}
	return nil
}
