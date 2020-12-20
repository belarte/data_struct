package list_test

import "testing"
import "github.com/belarte/data_struct/list"

func TestSimpleComparerCompare(t *testing.T) {
	tests := []struct {
		left     int
		right    int
		expected bool
	}{
		{0, 0, true},
		{0, 1, false},
		{1, 0, false},
		{1, 1, true},
		{42, 42, true},
		{42, 86, false},
		{86, 42, false},
		{86, 86, true},
	}

	var comparer list.Comparer = &list.SimpleComparer{}
	for _, test := range tests {
		got := comparer.Compare(test.left, test.right)
		if got != test.expected {
			t.Errorf("%v == %v is %v", test.left, test.right, got)
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

	var comparer list.Comparer = &list.SimpleComparer{}
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
