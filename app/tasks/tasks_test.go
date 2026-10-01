package tasks

import "testing"

func TestTitle(t *testing.T) {
	if got := Title(Task{ID: 7, Name: "Ship the merge queue"}); got != "#7 Ship the merge queue" {
		t.Errorf("Title() = %q", got)
	}
}
