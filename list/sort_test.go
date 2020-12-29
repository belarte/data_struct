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

func TestSimpleComparerCompare(t *testing.T) {
	tests := []struct {
		left     int
		right    int
		expected bool
	}{
		{0, 0, false},
		{0, 1, true},
		{1, 0, false},
		{1, 1, false},
		{42, 42, false},
		{42, 86, true},
		{86, 42, false},
		{86, 86, false},
	}

	var comparer list.Comparer = &list.LessThan{}
	for _, test := range tests {
		got := comparer.Compare(test.left, test.right)
		assert.Equal(t, test.expected, got, "%v < %v is %v", test.left, test.right, got)
	}
}

func TestSimpleComparerCount(t *testing.T) {
	tests := []struct {
		call  bool
		count int
	}{
		{false, 0},
		{true, 1},
		{false, 1},
		{true, 2},
		{true, 3},
		{true, 4},
		{false, 4},
	}

	var comparer list.Comparer = &list.LessThan{}
	for _, test := range tests {
		if test.call {
			comparer.Compare(0, 0)
		}
		got := comparer.Count()
		assert.Equal(t, test.count, got, "Expected count should be %v, but is %v", test.count, got)
	}
}

func TestSimpleSwapperSwap(t *testing.T) {
	tests := []struct {
		left          int
		right         int
		expectedLeft  int
		expectedRight int
	}{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{42, 86, 86, 42},
	}

	var swapper list.Swapper = &list.SimpleSwapper{}
	for _, test := range tests {
		swapper.Swap(&test.left, &test.right)
		assert.Equal(t, test.expectedLeft, test.left, "Left is %v but should be %v", test.left, test.expectedLeft)
		assert.Equal(t, test.expectedRight, test.right, "Right is %v but should be %v", test.right, test.expectedRight)
	}
}

func TestSimpleSwapperCount(t *testing.T) {
	tests := []struct {
		call  bool
		count int
	}{
		{false, 0},
		{true, 1},
		{false, 1},
		{true, 2},
		{true, 3},
		{true, 4},
		{false, 4},
	}

	var swapper list.Swapper = &list.SimpleSwapper{}
	for _, test := range tests {
		if test.call {
			left, right := 0, 0
			swapper.Swap(&left, &right)
		}
		got := swapper.Count()
		assert.Equal(t, test.count, got, "Expected count should be %v, but is %v", test.count, got)
	}
}

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
		var sorter list.Sorter = &list.SelectionSorter{&list.LessThan{}, swapper}
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, swapper.Count(), test.name)
	}
}

func TestSelectionsorterCallsSwapper(t *testing.T) {
	swapper := &mocks.Swapper{}
	defer swapper.AssertExpectations(t)
	var sorter list.Sorter = &list.SelectionSorter{&list.LessThan{}, swapper}
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
		var sorter list.Sorter = &list.SelectionSorter{comparer, &list.SimpleSwapper{}}
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, comparer.Count(), test.name)
	}
}

func TestSelectionsorterCallsComparer(t *testing.T) {
	comparer := &mocks.Comparer{}
	defer comparer.AssertExpectations(t)
	var sorter list.Sorter = &list.SelectionSorter{comparer, &list.SimpleSwapper{}}
	comparer.On("Compare", mock.Anything, mock.Anything).Times(10).Return(false)
	sorter.Sort([]int{3, 1, 5, 4, 2})
}

func TestSelectionSorter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element list", []int{42}, []int{42}},
		{"sorted pair", []int{2, 4}, []int{2, 4}},
		{"reverse sorted list", []int{4, 2}, []int{2, 4}},
		{"sorted list", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted list", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"shuffled list", []int{2, 5, 3, 6, 1, 4}, []int{1, 2, 3, 4, 5, 6}},
	}

	var sorter list.Sorter = &list.SelectionSorter{&list.LessThan{}, &list.SimpleSwapper{}}
	for _, test := range tests {
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, test.input, test.name)
	}
}

func TestInsertionSorter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element list", []int{42}, []int{42}},
		{"sorted pair", []int{2, 4}, []int{2, 4}},
		{"reverse sorted list", []int{4, 2}, []int{2, 4}},
		{"sorted list", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted list", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"shuffled list", []int{2, 5, 3, 6, 1, 4}, []int{1, 2, 3, 4, 5, 6}},
	}

	var sorter list.Sorter = &list.InsertionSorter{}
	for _, test := range tests {
		sorter.Sort(test.input)
		assert.Equal(t, test.expected, test.input, test.name)
	}
}

func BenchmarkSelectionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = &list.SelectionSorter{&list.LessThan{}, &list.SimpleSwapper{}}
		sorter.Sort(rand.Perm(2048))
	}
}

func BenchmarkInsertionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = &list.InsertionSorter{}
		sorter.Sort(rand.Perm(2048))
	}
}
