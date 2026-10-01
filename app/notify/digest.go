// Package notify builds the messages people get about their tasks.
package notify

import (
	"strings"

	"github.com/mergequeue-sandbox/sandbox/app/tasks"
)

// Digest lists the open tasks, one per line, for the daily email.
func Digest(ts []tasks.Task) string {
	var b strings.Builder
	for _, t := range ts {
		if !t.Done {
			b.WriteString("• " + tasks.DisplayTitle(t) + "\n")
		}
	}
	return b.String()
}
