package main

import (
	"fmt"
	"github.com/study/golang/day05/test/utils"
)

func main() {
	a, b := 2, 3
	fmt.Printf("a + b = %d\n", utils.Add(a, b))
	fmt.Printf("a - b = %d\n", utils.Sub(a, b))
	fmt.Printf("a * b = %d\n", utils.Mul(a, b))
	fmt.Printf("a / b = %d\n", utils.Div(a, b))
}