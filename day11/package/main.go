package main

import (
	"fmt"
	"time"
	_ "math/rand" // import the package, but does not use any item in the package
				  // the init function will been called

	"github.com/study/golang/day11/util"
)

const (
	N = 5000
)

func show(a []int) {
	for _, item := range a {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}

var a int = 1000 // exec before init
var b int = 2000 // exec before init

func init() {
	fmt.Printf("main init, a=[%d], b=[%d]\n", a, b)
}

//there can have mutil init-function
func init() {
	fmt.Printf("main init 2\n")
}

func main() {
	var array [N]int
	for i := 0; i < N; i++ {
		array[i] = N - i
	}

	show(array[:10])
	t1 := time.Now().UnixNano()
	util.SelectSort(array[:])
	t2 := time.Now().UnixNano()
	show(array[:10])

	fmt.Printf("sort time : %d micro seconds\n", (t2-t1)/1000/1000)
}
