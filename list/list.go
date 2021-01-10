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

// Assigner copy a value from source to destination
type Assigner interface {
	Assign(destination *int, source int)
	Count() int
}

// SimpleAssigner copy a value from source to destination and count number of calls
type SimpleAssigner struct {
	count int
}

// Assign assigns the value from source to destination
func (a *SimpleAssigner) Assign(destination *int, source int) {
	*destination = source
	a.count++
}

// Count returns the number of time Compare has been called
func (a *SimpleAssigner) Count() int {
	return a.count
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
