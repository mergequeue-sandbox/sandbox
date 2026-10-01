// Package projects groups tasks.
package projects

import "github.com/mergequeue-sandbox/sandbox/app/config"

// Page returns the bounds of page n (0-based) of a list of total items.
func Page(n, total int) (from, to int) {
	from = min(n*config.PageSize, total)
	return from, min(from+config.PageSize, total)
}
