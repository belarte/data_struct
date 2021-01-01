package main

import (
	"fmt"
	"github.com/belarte/data_struct/list"
	"math/rand"
)

func run(name string, count int) {
	fmt.Printf("Running %v sorter with list of %v elements\n", name, count)
	l := rand.Perm(count)
	swapper := &list.SimpleSwapper{}
	comparer := &list.LessThan{}
	sorter := list.NewSorter(name, comparer, swapper)
	sorter.Sort(l)
	fmt.Println("Swap.Count    ", swapper.Count())
	fmt.Println("Compare.Count ", comparer.Count())
}

func main() {
	run("insertion", 2048)
	run("selection", 2048)
}
