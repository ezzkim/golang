package main

import (
	"fmt"
	"time"

	"github.com/study/golang/day03/sort/util"
)

func main() {
	testSort()
}

const (
	Number = 50000
)

func testSort() {
	var array [Number]int
	for i := 0; i < Number; i++ {
		array[i] = Number - i
	}
	show(array[:10])

	start := time.Now().UnixNano()
	//util.SelectSort(array[:])
	//util.BubbleSort(array[:])
	util.InsertSort(array[:])
	end := time.Now().UnixNano()

	cost := (end - start) / 1000
	show(array[:10])

	fmt.Printf("cost time %d micro seconds\n", cost)
}

func show(array []int) {
	for _, item := range array {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}
