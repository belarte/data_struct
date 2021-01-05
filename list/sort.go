package list

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
