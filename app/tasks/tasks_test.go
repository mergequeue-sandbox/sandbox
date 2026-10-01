package tasks

import "testing"

func TestDisplayTitle(t *testing.T) {
	for _, c := range []struct {
		task Task
		want string
	}{
		{Task{ID: 7, Name: "Ship the merge queue"}, "#7 Ship the merge queue"},
		{Task{ID: 8, Name: "Write the design doc", Done: true}, "✓ #8 Write the design doc"},
	} {
		if got := DisplayTitle(c.task); got != c.want {
			t.Errorf("DisplayTitle(%+v) = %q, want %q", c.task, got, c.want)
		}
	}
}
