package list_test

import (
	"testing"

	"github.com/belarte/data_struct/list"
	"github.com/stretchr/testify/assert"
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

func TestSimpleAssignerAssign(t *testing.T) {
	tests := []struct {
		destination int
		source      int
	}{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
		{42, 42},
		{42, 86},
		{86, 42},
		{86, 86},
	}

	var assigner list.Assigner = &list.SimpleAssigner{}
	for _, test := range tests {
		assigner.Assign(&test.destination, test.source)
		assert.Equal(t, test.destination, test.source, "%v != %v", test.destination, test.source)
	}
}

func TestSimpleAssignerCount(t *testing.T) {
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

	var assigner list.Assigner = &list.SimpleAssigner{}
	for _, test := range tests {
		if test.call {
			temp := 0
			assigner.Assign(&temp, 0)
		}
		got := assigner.Count()
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
