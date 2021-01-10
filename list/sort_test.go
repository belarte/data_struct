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
		var sorter list.Sorter = list.NewSorter(
			"selection",
			&list.LessThan{},
			&list.SimpleAssigner{},
			swapper,
			&list.NoPrint{})
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, swapper.Count(), test.name)
	}
}

func TestSelectionsorterCallsSwapper(t *testing.T) {
	swapper := &mocks.Swapper{}
	defer swapper.AssertExpectations(t)
	var sorter list.Sorter = list.NewSorter(
		"selection",
		&list.LessThan{},
		&list.SimpleAssigner{},
		swapper,
		&list.NoPrint{})
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
		var sorter list.Sorter = list.NewSorter(
			"selection",
			comparer,
			&list.SimpleAssigner{},
			&list.SimpleSwapper{},
			&list.NoPrint{})
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, comparer.Count(), test.name)
	}
}

func TestSelectionSorterCallsComparer(t *testing.T) {
	comparer := &mocks.Comparer{}
	defer comparer.AssertExpectations(t)
	var sorter list.Sorter = list.NewSorter(
		"selection",
		comparer,
		&list.SimpleAssigner{},
		&list.SimpleSwapper{},
		&list.NoPrint{})
	comparer.On("Compare", mock.Anything, mock.Anything).Times(10).Return(false)
	sorter.Sort([]int{3, 1, 5, 4, 2})
}

func TestSelectionSorterNeverCallsAssigner(t *testing.T) {
	assigner := &mocks.Assigner{}
	defer assigner.AssertExpectations(t)
	var sorter list.Sorter = list.NewSorter(
		"selection",
		&list.LessThan{},
		assigner,
		&list.SimpleSwapper{},
		&list.NoPrint{})
	assigner.AssertNotCalled(t, "Assign")
	sorter.Sort([]int{3, 1, 5, 4, 2})
}

func testSorters(t *testing.T, name string) {
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

	var sorter list.Sorter = list.NewSorter(
		name,
		&list.LessThan{},
		&list.SimpleAssigner{},
		&list.SimpleSwapper{},
		&list.NoPrint{})
	for _, test := range tests {
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, test.input, test.name)
	}
}

func TestSelectionSorter(t *testing.T) {
	testSorters(t, "selection")
}

func TestInsertionSorter(t *testing.T) {
	testSorters(t, "insertion")
}

func TestMergeSorter(t *testing.T) {
	testSorters(t, "merge")
}

func TestParallelMergeSorter(t *testing.T) {
	testSorters(t, "parallel_merge")
}

func benchmarkRunner(algo string, size int) {
	var sorter list.Sorter = list.NewSorter(
		algo,
		&list.LessThan{},
		&list.SimpleAssigner{},
		&list.SimpleSwapper{},
		&list.NoPrint{})
	sorter.Sort(rand.Perm(size))
}

func BenchmarkSelectionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		benchmarkRunner("selection", 2048)
	}
}

func BenchmarkInsertionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		benchmarkRunner("insertion", 2048)
	}
}

func BenchmarkMergeSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		benchmarkRunner("merge", 2048)
	}
}

func BenchmarkParallelMergeSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		benchmarkRunner("parallel_merge", 2048)
	}
}
