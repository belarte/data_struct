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
	case "merge":
		return &mergeSorter{c, s, p}
	default:
		return nil
	}
}

type selectionSorter struct {
	comparer Comparer
	swapper  Swapper
	printer  Printer
}

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

type insertionSorter struct {
	comparer Comparer
	swapper  Swapper
	printer  Printer
}

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

type mergeSorter struct {
	comparer Comparer
	swapper  Swapper
	printer  Printer
}

func (s mergeSorter) Sort(list List) {
	s.printer.Print(list)
	s.sort(list, 0, len(list))
}

func (s mergeSorter) sort(list List, start, end int) {
	if end-start < 2 {
		return
	}

	mid := start + (end-start)/2
	s.sort(list, start, mid)
	s.sort(list, mid, end)

	for start < mid {
		if s.comparer.Compare(list[mid], list[start]) {
			s.swapper.Swap(&list[start], &list[mid])
			s.printer.Print(list)
			for i := mid; i < end-1 && s.comparer.Compare(list[i+1], list[i]); i++ {
				s.swapper.Swap(&list[i], &list[i+1])
				s.printer.Print(list)
			}
		}
		start++
	}
}
