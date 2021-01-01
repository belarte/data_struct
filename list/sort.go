package list

// List to be sorted
type List []int

// Comparer compares two ints
type Comparer interface {
	Compare(left, right int) bool
	Count() int
}

// LessThan checks for equality and counts number of calls
type LessThan struct {
	count int
}

// Compare returns true if the ints are equal
func (comp *LessThan) Compare(left, right int) bool {
	comp.count++
	return left < right
}

// Count returns the number of time Compare has been called
func (comp *LessThan) Count() int {
	return comp.count
}

// Swapper swaps two ints
type Swapper interface {
	Swap(left, right *int)
	Count() int
}

// SimpleSwapper inverts two ints and counts number of calls
type SimpleSwapper struct {
	count int
}

// Swap swaps two ints in place and counts number of calls
func (swap *SimpleSwapper) Swap(left, right *int) {
	swap.count++
	*left, *right = *right, *left
}

// Count returns the number of time Swap has been called
func (swap *SimpleSwapper) Count() int {
	return swap.count
}

// Sorter sorts list in place
type Sorter interface {
	Sort(List)
}

func NewSorter(name string, c Comparer, s Swapper) Sorter {
	switch name {
	case "selection":
		return &SelectionSorter{c, s}
	case "insertion":
		return &InsertionSorter{c, s}
	default:
		return nil

	}
}

// SelectionSorter implements selection sort on a list of integers
type SelectionSorter struct {
	comparer Comparer
	swapper  Swapper
}

// Sort sorts the list in place
func (s SelectionSorter) Sort(list List) {
	for i := 0; i < len(list)-1; i++ {
		index := i
		for j := i + 1; j < len(list); j++ {
			if s.comparer.Compare(list[j], list[index]) {
				index = j
			}
		}
		s.swapper.Swap(&list[i], &list[index])
	}
}

// InsertionSorter implements insertion sort on a list of integers
type InsertionSorter struct {
	comparer Comparer
	swapper  Swapper
}

// Sort sorts the list in place
func (s InsertionSorter) Sort(list List) {
	for i := 0; i < len(list); i++ {
		j := i
		for j > 0 && s.comparer.Compare(list[j], list[j-1]) {
			s.swapper.Swap(&list[j], &list[j-1])
			j--
		}
	}
}
