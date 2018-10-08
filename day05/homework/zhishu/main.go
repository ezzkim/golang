package main

import (
	"fmt"
)

func main() {
	array := make([]int, 0)
	for i := 1; i < 101; i++ {
		if isZhiShu(i) {
			array = append(array, i)
		}
	}

	fmt.Println(array)
}

func isZhiShu(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
