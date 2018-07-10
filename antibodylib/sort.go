package antibodylib

import (
	"sort"
	"strings"
)

type indexedLine struct {
	idx  int
	line string
}

type indexedLines []indexedLine

// Len is needed by Sort interface
func (slice indexedLines) Len() int {
	return len(slice)
}

// Less is needed by Sort interface
func (slice indexedLines) Less(i, j int) bool {
	return slice[i].idx < slice[j].idx
}

// Swap is needed by Sort interface
func (slice indexedLines) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Sort all lines and join them in a string
func (slice indexedLines) String() string {
	sort.Sort(slice)
	// nolint: prealloc
	var lines []string
	for _, line := range slice {
		lines = append(lines, line.line)
	}
	return strings.Join(lines, "\n")
}
