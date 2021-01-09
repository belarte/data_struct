package list_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/belarte/data_struct/list"
	"github.com/belarte/data_struct/list/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectionSorterCallsSwapperAndReturnsCorrectCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty list", []int{}, 0},
		{"single element list", []int{42}, 0},
		{"sorted pair", []int{2, 4}, 1},
		{"reverse sorted list", []int{4, 2}, 1},
		{"sorted list", []int{1, 2, 3, 4, 5}, 4},
		{"reverse sorted list", []int{5, 4, 3, 2, 1}, 4},
		{"shuffled list", []int{2, 5, 3, 6, 1, 4}, 5},
	}

	for _, test := range tests {
		var swapper list.Swapper = &list.SimpleSwapper{}
		var sorter list.Sorter = list.NewSorter("selection", &list.LessThan{}, swapper, &list.NoPrint{})
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, swapper.Count(), test.name)
	}
}

func TestSelectionsorterCallsSwapper(t *testing.T) {
	swapper := &mocks.Swapper{}
	defer swapper.AssertExpectations(t)
	var sorter list.Sorter = list.NewSorter("selection", &list.LessThan{}, swapper, &list.NoPrint{})
	swapper.On("Swap", mock.Anything, mock.Anything).Times(4)
	sorter.Sort([]int{3, 1, 5, 4, 2})
}

func TestSelectionSorterCallsComparerAndReturnsCorrectCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty list", []int{}, 0},
		{"single element list", []int{42}, 0},
		{"sorted pair", []int{2, 4}, 1},
		{"reverse sorted list", []int{4, 2}, 1},
		{"sorted list", []int{1, 2, 3, 4, 5}, 10},
		{"reverse sorted list", []int{5, 4, 3, 2, 1}, 10},
		{"shuffled list", []int{2, 5, 3, 6, 1, 4}, 15},
	}

	for _, test := range tests {
		var comparer list.Comparer = &list.LessThan{}
		var sorter list.Sorter = list.NewSorter("selection", comparer, &list.SimpleSwapper{}, &list.NoPrint{})
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, comparer.Count(), test.name)
	}
}

func TestSelectionsorterCallsComparer(t *testing.T) {
	comparer := &mocks.Comparer{}
	defer comparer.AssertExpectations(t)
	var sorter list.Sorter = list.NewSorter("selection", comparer, &list.SimpleSwapper{}, &list.NoPrint{})
	comparer.On("Compare", mock.Anything, mock.Anything).Times(10).Return(false)
	sorter.Sort([]int{3, 1, 5, 4, 2})
}

func testSorters(t *testing.T, sorter list.Sorter) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element list", []int{42}, []int{42}},
		{"sorted pair", []int{2, 4}, []int{2, 4}},
		{"reverse sorted pair", []int{4, 2}, []int{2, 4}},
		{"sorted list", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted list", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"shuffled list", []int{2, 5, 3, 6, 1, 4}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range tests {
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, test.input, test.name)
	}
}

func TestSelectionSorter(t *testing.T) {
	var sorter list.Sorter = list.NewSorter("selection", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
	testSorters(t, sorter)
}

func TestInsertionSorter(t *testing.T) {
	var sorter list.Sorter = list.NewSorter("insertion", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
	testSorters(t, sorter)
}

func TestMergeSorter(t *testing.T) {
	var sorter list.Sorter = list.NewSorter("merge", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
	testSorters(t, sorter)
}

func TestParallelMergeSorter(t *testing.T) {
	var sorter list.Sorter = list.NewSorter("parallel_merge", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
	testSorters(t, sorter)
}

func BenchmarkSelectionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = list.NewSorter("selection", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
		sorter.Sort(rand.Perm(2048))
	}
}

func BenchmarkInsertionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = list.NewSorter("insertion", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
		sorter.Sort(rand.Perm(2048))
	}
}

func BenchmarkMergeSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = list.NewSorter("merge", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
		sorter.Sort(rand.Perm(2048))
	}
}

func BenchmarkParallelMergeSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = list.NewSorter("parallel_merge", &list.LessThan{}, &list.SimpleSwapper{}, &list.NoPrint{})
		sorter.Sort(rand.Perm(2048))
	}
}
