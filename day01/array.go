package main

import (
	"fmt"
)

func sum(a []int) int {
	var s int = 0
	for _, v := range a {
		s += v
	}

	return s
}

func max(a []int) int {
	fmt.Println("content = ", a, ", len = ", len(a), ", cap = ", cap(a))
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}

	return m
}

func through() {
	sti := 2
	switch {
	case sti >= 0:
		fmt.Println("sti>=0")
		//break, default
		fallthrough
	case sti >= 1:
		fmt.Println("sti>=1")
	case sti >= 2:
		fmt.Println("sti>=2")
	default:
		fmt.Println("default")
	}
}

func array() {
	var a [2]int
	var b [1]int = [1]int{11}
	c := [5]int{1, 2, 3, 4, 5}

	d := [...]int{5: 1, 2, 10}
	fmt.Println(a, b, c, d)
}

func arrayTest() {
	fmt.Println("---array test begin---")

	s := []int{1, 2, 3, 4, 5}
	su := sum(s)
	fmt.Printf("sum = %d\n", su)

	m := max(s)
	fmt.Printf("max = %d\n", m)

	var ss string = "ericsson"
	for i, c := range ss {
		fmt.Printf("%d -> %c", i, c)
	}
	fmt.Println("")

	through()

	fmt.Println("---begin array---")
	array()
	fmt.Println("---end array---")

	fmt.Println("---array test end---")
}
