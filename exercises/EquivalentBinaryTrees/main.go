
package main

import "golang.org/x/tour/tree"
import (
	"fmt"
	"sort"
	"time"
)

const TREESIZE = 10
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree) []int {
	var slc []int = make([]int, 0, TREESIZE)
	
	walk_recursive(t, &slc)
	
	return slc
}
func walk_recursive(t *tree.Tree, slc *[]int) {
	if (t.Left != nil) {
		walk_recursive(t.Left, slc)
	}
	if (t.Right != nil) {
		walk_recursive(t.Right, slc)
	}
	*slc = append(*slc, t.Value)
}

func walkTreeSorted(t *tree.Tree, ch chan int) {
	var slc []int = Walk(t)
	sort.Sort(sort.IntSlice(slc))
	for _, v := range slc {
		ch <- v
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree, ch chan bool) {
	ch1 := make(chan int, TREESIZE)
	ch2 := make(chan int, TREESIZE)
	go walkTreeSorted(t1, ch1)
	go walkTreeSorted(t2, ch2)
	
	for i := 0; i < TREESIZE; i++ {
		if <-ch1 != <-ch2 {
			ch <- false
			return
		}
	}
	
	ch <- true
	return
}

func Same_single(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, TREESIZE)
	ch2 := make(chan int, TREESIZE)
	walkTreeSorted(t1, ch1)
	walkTreeSorted(t2, ch2)
	
	for i := 0; i < TREESIZE; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	
	return true
}

func main() {
	start := time.Now()
	
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	
	go Same(tree.New(1), tree.New(1), ch1)
	go Same(tree.New(1), tree.New(2), ch2)
	
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	
	t := time.Now()
	fmt.Printf("multithreaded: %v.\n", t.Sub(start))


	start = time.Now()
	
	fmt.Println(Same_single(tree.New(1), tree.New(1)))
	fmt.Println(Same_single(tree.New(1), tree.New(2)))
	
	t = time.Now()
	fmt.Printf("singlethreaded: %v.\n", t.Sub(start))

	/*
		the non-multithreaded solution is consistently significantly faster, and its still using channels for the Walk() function. Maybe that's just because of the decreased overhead for this small example, but I wonder if I'm not doing this optimally.
	*/
}
