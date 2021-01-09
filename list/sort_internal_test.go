package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSorterMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"reverse list", []int{5, 4}, []int{4, 5}},
		{"reverse list", []int{2, 1}, []int{1, 2}},
		{"reverse list", []int{3, 1, 2}, []int{1, 2, 3}},
		{"sorted list", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"reverse list", []int{5, 6, 7, 8, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"interleaved 1 list", []int{1, 3, 5, 7, 2, 4, 6, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"interleaved 2 list", []int{2, 4, 6, 8, 1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	sorter := mergeSorter{&LessThan{}, &SimpleSwapper{}, &NoPrint{}}
	for _, test := range tests {
		sorter.merge(test.input, 0, len(test.input)/2, len(test.input))
		assert.Equal(t, test.expected, test.input, test.name)
	}
}
