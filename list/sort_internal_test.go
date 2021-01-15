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
		{"reverse pair", []int{5, 4}, []int{4, 5}},
		{"reverse list", []int{2, 1}, []int{1, 2}},
		{"reverse list", []int{3, 1, 2}, []int{1, 2, 3}},
		{"sorted list", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"reverse list", []int{5, 6, 7, 8, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"interleaved 1 list", []int{1, 3, 5, 7, 2, 4, 6, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"interleaved 2 list", []int{2, 4, 6, 8, 1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	sorter := mergeSorter{&LessThan{}, &SimpleAssigner{}, &SimpleSwapper{}, &NoPrint{}}
	for _, test := range tests {
		sorter.merge(test.input, 0, len(test.input)/2, len(test.input))
		assert.Equal(t, test.expected, test.input, test.name)
	}
}

func TestQuickSorterPivot(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		value    int
		left     []int
		right    []int
	}{
		{"sorted pair", []int{4, 5}, 1, 5, []int{4}, []int{5}},
		{"reverse pair", []int{5, 4}, 0, 4, []int{}, []int{4, 5}},
		{"sorted list", []int{1, 2, 3, 4, 5, 6, 7, 8}, 7, 8, []int{1, 2, 3, 4, 5, 6, 7}, []int{8}},
		{"reverse list", []int{8, 7, 6, 5, 4, 3, 2, 1}, 0, 1, []int{}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"interleaved 1 list", []int{1, 3, 5, 7, 2, 4, 6, 8, 5}, 4, 5, []int{1, 2, 3, 4}, []int{5, 5, 6, 7, 8}},
		{"interleaved 2 list", []int{2, 4, 6, 8, 1, 3, 5, 7}, 6, 7, []int{1, 2, 3, 4, 5, 6}, []int{7, 8}},
	}

	sorter := quickSorter{&LessThan{}, &SimpleAssigner{}, &SimpleSwapper{}, &NoPrint{}}
	for _, test := range tests {
		got := sorter.pivot(test.input, 0, len(test.input))
		val := test.input[got]
		left := test.input[0:got]
		right := test.input[got:len(test.input)]
		assert.Equalf(t, test.expected, got, "%s: pivot index is wrong", test.name)
		assert.Equalf(t, test.value, val, "%s: pivot value is wrong", test.name)
		assert.ElementsMatch(t, test.left, left, test.name)
		assert.ElementsMatch(t, test.right, right, test.name)
	}
}
