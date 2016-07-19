// Package sortstr provides sorting methods for string
// slices based on multiple indices or column headers.
// Reverse sort is possible by providing negative index
// values or prefixing column header titles with "-".
// Note that indices are not zero based, but start with 1.
// If an index is out of range, an empty string will be
// used for comparing, rather than throwing a runtime
// panic.
//
// Normally you only need to use this function:
//
//     By(rows, indices)
//
// or these two together:
//
//     NewHeaders(titles)
//     ByHeaders(headers, rows, titles)
//
// Only if you need to sort the same rows multiple times
// in different orders, you'll need to use the Multi type
// directly.
//
// This was developed with
// http://godoc.org/github.com/tealeg/xlsx
// in mind, but can be used independently.
package sortstr

import (
	"fmt"
	"sort"
	"strings"
)

// get retrieves value by column index, returns empty
// string if doesn't exist.
func get(row []string, col int) string {
	if col < len(row) {
		return row[col]
	}
	return ""
}

// Headers defines a map of column indices (int) by header
// title (string).
type Headers map[string]int

// NewHeaders creates new Headers from column header
// titles.
func NewHeaders(titles []string) Headers {
	hs := make(Headers)
	for i, title := range titles {
		if title != "" {
			hs[title] = i + 1
			hs[fmt.Sprintf("-%s", title)] = -hs[title]
		}
	}
	return hs
}

// Index of a given column header title
func (hs Headers) Index(title string) (int, error) {
	if i, ok := hs[title]; ok {
		return i, nil
	}
	return 0, fmt.Errorf("Unknown column header: %s (%#v)",
		title, hs)
}

// Indices of given column header titles
func (hs Headers) Indices(titles ...string) (
	[]int, error) {
	indices := make([]int, len(titles))
	for i, title := range titles {
		index, err := hs.Index(title)
		if err != nil {
			return []int{}, err
		}
		indices[i] = index
	}
	return indices, nil
}

// Multi implements the Sort interface, sorting the rows
// within, based on multiple column indices. For sorting
// rows you would normally use the shortcut functions:
//
//     sortstr.By
//     sortstr.ByHeaders
type Multi struct {
	Rows    [][]string
	Indices []int
}

// By sorts rows in place by using the indices. If indices
// are negative, they will be sorted in reverse order.
func By(rows [][]string, indices ...int) {
	(&Multi{
		Rows:    rows,
		Indices: indices,
	}).Sort()
}

// ByHeaders sorts rows in place by using the column
// header titles. If titles are prefixed with "-", they
// will be sorted in reverse order.
func ByHeaders(headers Headers, rows [][]string,
	titles ...string) error {
	indices, err := headers.Indices(titles...)
	if err != nil {
		return err
	}
	(&Multi{
		Rows:    rows,
		Indices: indices,
	}).Sort()
	return nil
}

// Sort sorts the argument slice according to its indices
func (m *Multi) Sort() {
	sort.Sort(m)
}

// Len is part of sort.Interface.
func (m *Multi) Len() int {
	return len(m.Rows)
}

// Swap is part of sort.Interface.
func (m *Multi) Swap(i, j int) {
	m.Rows[i], m.Rows[j] = m.Rows[j], m.Rows[i]
}

// getReverse returns a positive column index and a bool to
// indicate the order should be reserved
func getReverse(index int) (int, bool) {
	reverse := false
	if index < 0 {
		reverse = true
		index = -index
	}
	return index - 1, reverse
}

// Less is part of sort.Interface. It is implemented by
// looping along the indices until it finds a comparison
// that is either Less or !Less. Note that it can call the
// indices comparisons twice per call.
func (m *Multi) Less(i, j int) bool {
	p, q := m.Rows[i], m.Rows[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(m.Indices)-1; k++ {
		index, reverse := getReverse(m.Indices[k])
		switch {
		case get(p, index) < get(q, index):
			// p < q, so we have a decision.
			return !reverse
		case get(q, index) < get(p, index):
			// p > q, so we have a decision.
			return reverse
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return
	// whatever the final comparison reports.
	index, reverse := getReverse(m.Indices[k])
	if get(p, index) < get(q, index) {
		return !reverse
	}
	return reverse
}

// Print rows (mostly for debugging purposes). See the
// examples of By and ByHeaders.
func Print(label string, rows [][]string, sep string) {
	fmt.Printf("\n%s:\n", label)
	for _, row := range rows {
		fmt.Println(strings.Join(row, sep))
	}
}
