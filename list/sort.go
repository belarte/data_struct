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

// Swaper swaps two ints
type Swapper interface {
	Swap(left, right *int)
	Count() int
}

// SimpleSwaper inverts two ints and counts number of calls
type SimpleSwapper struct {
	count int
}

func (swap *SimpleSwapper) Swap(left, right *int) {
	swap.count = swap.count + 1
	*left, *right = *right, *left
}

func (swap *SimpleSwapper) Count() int {
	return swap.count
}

// Sorter sorts list in place
type Sorter interface {
	Sort(List)
}
