package notify

import (
	"fmt"

	"github.com/mergequeue-sandbox/sandbox/app/tasks"
)

// Weekly is the Monday summary: how many tasks were done, and the first
// one still open.
func Weekly(ts []tasks.Task) string {
	done := 0
	var next string
	for _, t := range ts {
		if t.Done {
			done++
		} else if next == "" {
			next = tasks.DisplayTitle(t)
		}
	}
	return fmt.Sprintf("%d done this week. Next up: %s", done, next)
}
