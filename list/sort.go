package list

// List to be sorted
type List []int

// Comparer compares two ints
type Comparer interface {
	Compare(left, right int) bool
	Count() int
}

// SimpleComparer checks for equality and counts number of calls
type SimpleComparer struct {
	count int
}

// Compare returns true if the ints are equal
func (comp *SimpleComparer) Compare(left, right int) bool {
	comp.count = comp.count + 1
	return left == right
}

// Count returns the number of time Compare has been called
func (comp *SimpleComparer) Count() int {
	return comp.count
}

// Sorter sorts list in place
type Sorter interface {
	Sort(List)
}
