package main

import (
	"flag"
	"fmt"
	"github.com/belarte/data_struct/list"
	"math/rand"
)

func run(name string, size int) {
	fmt.Printf("Running %v sorter with list of size %v\n", name, size)
	l := rand.Perm(size)
	swapper := &list.SimpleSwapper{}
	comparer := &list.LessThan{}
	sorter := list.NewSorter(name, comparer, swapper, &list.NoPrint{})
	sorter.Sort(l)
	fmt.Println("Swap.Count    ", swapper.Count())
	fmt.Println("Compare.Count ", comparer.Count())
}

func main() {
	size := flag.Int("size", 8, "Number of elements per list.")
	algo := flag.String("algorithm", "", "Sorting algorith to use. If none specified, will run all algorithms.")
	flag.Parse()

	switch *algo {
	case "insertion", "selection":
		run(*algo, *size)
	case "":
		run("insertion", *size)
		run("selection", *size)
	default:
		fmt.Println("Unknown sorting algorithm...")
	}
}
