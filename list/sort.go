package list

type sortersList []string

// Sorters is a list of available sorting algorithms
var Sorters sortersList = []string{"insertion", "selection", "merge", "parallel_merge", "quick"}

// Contains returns true if the parameter is a valid sorting algorithm
func (s sortersList) Contains(name string) bool {
	for _, v := range s {
		if name == v {
			return true
		}
	}
	return false
}

// Sorter sorts list in place
type Sorter interface {
	Sort(List)
}

// NewSorter creates a new sorting algorithm
func NewSorter(name string, c Comparer, a Assigner, s Swapper, p Printer) Sorter {
	switch name {
	case "selection":
		return &selectionSorter{c, a, s, p}
	case "insertion":
		return &insertionSorter{c, a, s, p}
	case "merge":
		return &mergeSorter{c, a, s, p}
	case "parallel_merge":
		return &parallelMergeSorter{c, a, s, p}
	case "quick":
		return &quickSorter{c, a, s, p}
	default:
		return nil
	}
}

type selectionSorter struct {
	comparer Comparer
	assigner Assigner
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
	assigner Assigner
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
	assigner Assigner
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
	s.merge(list, start, mid, end)
	s.printer.Print(list)
}

func (s mergeSorter) merge(list List, start, mid, end int) {
	lower := make([]int, mid-start)
	copy(lower, list[start:mid])
	low := 0
	up := mid
	for i := start; i < end && low < mid-start && up < end; i++ {
		if s.comparer.Compare(lower[low], list[up]) {
			s.assigner.Assign(&list[i], lower[low])
			low++
		} else {
			s.assigner.Assign(&list[i], list[up])
			up++
		}
	}

	again := start + (end - mid)
	for low < mid-start {
		s.assigner.Assign(&list[again+low], lower[low])
		low++
	}

}

type parallelMergeSorter struct {
	comparer Comparer
	assigner Assigner
	swapper  Swapper
	printer  Printer
}

func (s parallelMergeSorter) Sort(list List) {
	s.printer.Print(list)
	s.sort(list, 0, len(list))
}

func (s parallelMergeSorter) sort(list List, start, end int) {
	if end-start < 2 {
		return
	}

	mid := start + (end-start)/2
	done := make(chan bool, 1)
	go func() {
		s.sort(list, start, mid)
		done <- true
	}()
	s.sort(list, mid, end)
	<-done
	s.merge(list, start, mid, end)
	s.printer.Print(list)
}

func (s parallelMergeSorter) merge(list List, start, mid, end int) {
	lower := make([]int, mid-start)
	copy(lower, list[start:mid])
	low := 0
	up := mid
	for i := start; i < end && low < mid-start && up < end; i++ {
		if s.comparer.Compare(lower[low], list[up]) {
			s.assigner.Assign(&list[i], lower[low])
			low++
		} else {
			s.assigner.Assign(&list[i], list[up])
			up++
		}
	}

	again := start + (end - mid)
	for low < mid-start {
		s.assigner.Assign(&list[again+low], lower[low])
		low++
	}

}

type quickSorter struct {
	comparer Comparer
	assigner Assigner
	swapper  Swapper
	printer  Printer
}

func (s quickSorter) Sort(list List) {
	s.printer.Print(list)
	s.sort(list, 0, len(list))
}

func (s quickSorter) sort(list List, start, end int) {
	if end-start < 2 {
		return
	}

	pivot := s.pivot(list, start, end)
	s.sort(list, start, pivot)
	s.sort(list, pivot+1, end)
}

func (s quickSorter) pivot(list List, start, end int) int {
	lastIndex := end - 1
	value := list[lastIndex]

	pivot := start
	for i := start; i < end; i++ {
		if s.comparer.Compare(list[i], value) {
			s.swapper.Swap(&list[pivot], &list[i])
			s.printer.Print(list)
			pivot++
		}
	}

	s.swapper.Swap(&list[pivot], &list[end-1])
	s.printer.Print(list)
	return pivot
}
