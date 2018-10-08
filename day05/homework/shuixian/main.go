package main

import (
	"fmt"
)

func main() {
	array := make([]int, 0)
	for i := 100; i < 1000; i++ {
		if isShuiXian(i) {
			array = append(array, i)
		}
	}

	fmt.Println(array)
}

func isShuiXian(n int) bool {
	if n < 100 || n > 999 {
		return false
	}

	g := n%10
	s := n/10%10
	b := n/100

	sum := g*g*g + s*s*s + b*b*b
	if sum == n {
		return true
	}

	return false
}
