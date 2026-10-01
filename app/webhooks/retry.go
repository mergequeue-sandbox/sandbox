package webhooks

import "time"

// Backoff is how long to wait before delivery attempt n (1-based): 10s,
// 20s, 40s, … capped at 10 minutes.
func Backoff(n int) time.Duration {
	d := 10 * time.Second << (n - 1)
	if d <= 0 || d > 10*time.Minute {
		return 10 * time.Minute
	}
	return d
}
