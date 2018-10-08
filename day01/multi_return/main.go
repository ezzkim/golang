package main

import (
	"fmt"
)

func add(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, sub := add(2, 5)
	fmt.Println("sum = ", sum, ", sub = ", sub)
	fmt.Printf("sum = %d, sub = %d\n", sum, sub)
}
