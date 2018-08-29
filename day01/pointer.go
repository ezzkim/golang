package main

import (
	"fmt"
)

func diffType() {
	var ip *int
	var fp *float32

	var aa int = 10
	var ff float32 = 32.5

	ip = &aa
	fp = &ff

	fmt.Printf("*ip = %d, *fp = %f\n", *ip, *fp)
}

func pointerTest() {
	fmt.Println("---begine pointer test---")
	diffType()
	fmt.Println("---end pointer test---")
}