package list_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/belarte/data_struct/list"
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
		if got != test.expected {
			t.Errorf("%v < %v is %v", test.left, test.right, got)
		}
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
		if got != test.count {
			t.Errorf("Expected count should be %v, but is %v", test.count, got)
		}
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
		if test.left != test.expectedLeft || test.right != test.expectedRight {
			t.Errorf("Swap(%v, %v) != (%v, %v)", test.left, test.right, test.expectedLeft, test.expectedRight)
		}
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
		if got != test.count {
			t.Errorf("Expected count should be %v, but is %v", test.count, got)
		}
	}
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

	var sorter list.Sorter = &list.SelectionSorter{}
	for _, test := range tests {
		sorter.Sort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("'%v' test failed, list is: %v but should be %v", test.name, test.input, test.expected)
		}
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
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("'%v' test failed, list is: %v but should be %v", test.name, test.input, test.expected)
		}
	}
}

func BenchmarkSelectionSorter(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		var sorter list.Sorter = &list.SelectionSorter{}
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
