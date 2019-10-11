package antibodylib

import (
	"sort"
	"strings"
	"sync"
)

type indexedLine struct {
	idx  int
	line string
}

type indexedLines []indexedLine

type safeIndexedLines struct {
	mutex sync.Mutex
	data  indexedLines
}

// Append safely appends items to the slice
func (slice *safeIndexedLines) Append(item indexedLine) {
	slice.mutex.Lock()
	defer slice.mutex.Unlock()

	slice.data = append(slice.data, item)
}

func (slice *safeIndexedLines) Items() indexedLines {
	slice.mutex.Lock()
	defer slice.mutex.Unlock()
	return slice.data
}

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
