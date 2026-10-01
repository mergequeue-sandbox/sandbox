package notify

import (
	"testing"

	"github.com/mergequeue-sandbox/sandbox/app/tasks"
)

func TestDigest(t *testing.T) {
	got := Digest([]tasks.Task{{ID: 1, Name: "Write docs"}, {ID: 2, Name: "Ship it", Done: true}})
	if want := "• #1 Write docs\n"; got != want {
		t.Errorf("Digest() = %q, want %q", got, want)
	}
}
