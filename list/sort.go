package list

import "fmt"

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

// Printer displays a list
type Printer interface {
	// Print displays the list
	Print(list List)
}

// NoPrint is to be used when no output is expected
type NoPrint struct{}

// Print does not do anything
func (p *NoPrint) Print(list List) {
}

// CommandLinePrinter is to be used when an output on the command line is desired
type CommandLinePrinter struct{}

// Print displays the list on the command line
func (p *CommandLinePrinter) Print(list List) {
	fmt.Println(list)
}

// Sorter sorts list in place
type Sorter interface {
	Sort(List)
}

// NewSorter creates a new sorting algorithm
func NewSorter(name string, c Comparer, s Swapper, p Printer) Sorter {
	switch name {
	case "selection":
		return &selectionSorter{c, s, p}
	case "insertion":
		return &insertionSorter{c, s, p}
	default:
		return nil

	}
}

// SelectionSorter implements selection sort on a list of integers
type selectionSorter struct {
	comparer Comparer
	swapper  Swapper
	printer  Printer
}

// Sort sorts the list in place
func (s selectionSorter) Sort(list List) {
	s.printer.Print(list)
	for i := 0; i < len(list)-1; i++ {
		index := i
		for j := i + 1; j < len(list); j++ {
			if s.comparer.Compare(list[j], list[index]) {
				index = j
			}
		}
		s.swapper.Swap(&list[i], &list[index])
		s.printer.Print(list)
	}
}

// InsertionSorter implements insertion sort on a list of integers
type insertionSorter struct {
	comparer Comparer
	swapper  Swapper
	printer  Printer
}

// Sort sorts the list in place
func (s insertionSorter) Sort(list List) {
	s.printer.Print(list)
	for i := 0; i < len(list); i++ {
		j := i
		for j > 0 && s.comparer.Compare(list[j], list[j-1]) {
			s.swapper.Swap(&list[j], &list[j-1])
			s.printer.Print(list)
			j--
		}
	}
}
