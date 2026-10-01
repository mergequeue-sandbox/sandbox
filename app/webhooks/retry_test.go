package webhooks

import (
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	for n, want := range map[int]time.Duration{1: 10 * time.Second, 3: 40 * time.Second, 20: 10 * time.Minute} {
		if got := Backoff(n); got != want {
			t.Errorf("Backoff(%d) = %v, want %v", n, got, want)
		}
	}
}
