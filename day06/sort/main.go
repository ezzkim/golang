package main

import (
	"fmt"
	"time"

	"github.com/study/golang/day06/sort/util"
)

const (
	N = 50000
)

func main() {
	var a [N]int
	for i := 0; i < N; i++ {
		a[i] = N - i
	}

	show(a[:10])
	t1 := time.Now().UnixNano()
	//util.SelectSort(a[:])
	//util.BubbleSort(a[:])
	util.InsertSort(a[:])
	t2 := time.Now().UnixNano()
	show(a[:10])
	fmt.Printf("t1=%d, t2=%d\n", t1, t2)
	fmt.Printf("sort use time = %d mico seconds\n", (t2-t1) / 1000 / 1000)
}

func show(a []int) {
	for _, item := range a {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}
