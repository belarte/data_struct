package main

import (
	"flag"
	"fmt"
	"github.com/belarte/data_struct/list"
	"math/rand"
	"time"
)

func run(name string, size int, printer list.Printer) {
	fmt.Printf("Running %v sorter with list of size %v\n", name, size)
	l := rand.Perm(size)
	swapper := &list.SimpleSwapper{}
	comparer := &list.LessThan{}
	sorter := list.NewSorter(name, comparer, swapper, printer)
	sorter.Sort(l)
	fmt.Println("Swap.Count    ", swapper.Count())
	fmt.Println("Compare.Count ", comparer.Count())
}

func main() {
	size := flag.Int("size", 8, "Number of elements per list.")
	algo := flag.String("algorithm", "", "Sorting algorith to use. If none specified, will run all algorithms.")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	switch *algo {
	case "insertion", "selection":
		run(*algo, *size, &list.CommandLinePrinter{})
	case "":
		run("insertion", *size, &list.NoPrint{})
		run("selection", *size, &list.NoPrint{})
	default:
		fmt.Println("Unknown sorting algorithm...")
	}
}
